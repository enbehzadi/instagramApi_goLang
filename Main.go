package main

import (
	"fmt"
	"gopkg.in/ahmdrz/goinsta.v2"
	"os"
	"time"
)

var instaLogin *goinsta.Instagram
func main() {

	//login
	instaLogin = login("mr.masoudi1234" , "multi123456")
	//searchLocation("36.2","59.6","")

	//searchuser("itsarezou")
	//follow("td110")
	searchHashtag("iran")
}
func follow(name string)  {
	if instaLogin != nil {

		user, e := instaLogin.Profiles.ByName(name)
		fmt.Print(user)
		//examples.CheckErr(err)

		if e == nil {

			fmt.Printf("Following: %v\n", user.Friendship.Following)
			e = user.Follow()
			fmt.Printf("After func call: Following: %v\n", user.Friendship.Following)
		}else {
			fmt.Print("name not found")
		}
	}else {
		fmt.Print("login fail")
	}
}

func searchLocation(lat string,lng string ,query string)  {
	if instaLogin != nil {

		rss,e:= instaLogin.Search.Location(lat, lng, query)
		fmt.Print(rss)
		//examples.CheckErr(err)

		if e == nil {


			// getting images URL
			for _, venue := range rss.Venues {
				fmt.Printf("    %s\n", venue.Name)
			}
		}else {
			fmt.Print("location not found")
		}
	}else {
		fmt.Print("login fail")
	}
}

func searchuser(name string)  {
	if instaLogin != nil {

		rss,e:= instaLogin.Search.User(name)
		fmt.Print(rss)
		//examples.CheckErr(err)

		if e == nil {


			// getting images URL
			for _, venue := range rss.Venues {
				fmt.Printf("    %s\n", venue.Name)
			}
		}else {
			fmt.Print("file not found")
		}
	}else {
		fmt.Print("login fail")
	}
}

func searchHashtag(Hashtag string)  {
	if instaLogin != nil {

		rss,e:= instaLogin.Search.Tags(Hashtag)
		fmt.Print(rss)
		//examples.CheckErr(err)

		if e == nil {


			// getting images URL
			for _, venue := range rss.Venues {
				fmt.Printf("    %s\n", venue.Name)
			}
		}else {
			fmt.Print("file not found")
		}
	}else {
		fmt.Print("login fail")
	}
}
func login(user , pass string) *goinsta.Instagram  {

	if instaLogin == nil {

		instaLogin = goinsta.New(user, pass )


		if err := instaLogin.Login(); err != nil {
			fmt.Println(err)
			fmt.Println("error")
			return  nil
		}else {
			fmt.Println("succed")

			//comment("2053608763040245760" ,"so nice")
			//like("2053608763040245760")
			//uploadPhoto("messi.jpg" , "#so beautiful 12")
			//getRecent()
			//setStories()
			return instaLogin

		}
	}else {

		return instaLogin
	}




}


func like(id string) string {

	if instaLogin != nil {

		media, err := instaLogin.GetMedia(id)
		fmt.Print(media)
		//examples.CheckErr(err)

		if err == nil {

			fmt.Printf("Liked: %v\n", media.Items[0].HasLiked)
			media.Items[0].Like()

			media.Sync()
			fmt.Printf("Liked: %v\n", media.Items[0].HasLiked)
			return "like done"

		}else {
			return "media not found"
		}


	}else {
		fmt.Print("login fail")
		return "login fail"
	}
}

func setStories(){

	//inst, err := e.InitGoinsta("<another user>")
	//e.CheckErr(err)

	user ,_ := instaLogin.Profiles.ByName(os.Args[0])


	stories := user.Stories()
	//e.CheckErr(err)

	for stories.Next() {
		// getting images URL
		for _, item := range stories.Items {
			if len(item.Images.Versions) > 0 {
				fmt.Printf("  Image - %s\n", item.Images.Versions[0].URL)
			}
			if len(item.Videos) > 0 {
				fmt.Printf("  Video - %s\n", item.Videos[0].URL)
			}
		}
	}
}

func uploadPhoto(path , caption string) string{

	if instaLogin != nil {
		file, err := os.Open(path)
		if (err == nil) {

			instaLogin.UploadPhoto(file, caption, 87, 0)
			return "upload done"
		} else {
			println("error uploading")
			return "error uploading"
		}

	}else {

		return "login failed"
	}

}




func comment(id , text string ){

	media, err := instaLogin.GetMedia(id)
	fmt.Print(media)
	//examples.CheckErr(err)
	if err == nil {


	}

	fmt.Printf("Comments: %d\n", media.Items[0].CommentCount)
	err = media.Items[0].Comments.Add(text)
	//examples.CheckErr(err)

	fmt.Println("wait 5 seconds...")
	for i := 5; i > 0; i-- {
		fmt.Printf("%d ", i)
		time.Sleep(time.Second)
	}
	fmt.Println()
	media.Sync()
	fmt.Printf("After calling: Comments: %d\n", media.Items[0].CommentCount)

	//tray, err := instaLogin.Timeline.Stories()
	////examples.CheckErr(err)
	//
	//story := tray.Stories[1]
	//// commenting your first timeline story xddxdxd
	//fmt.Printf("Sending reply to %s %s\n", story.Items[0].Images.GetBest(), story.Items[0].MediaToString())
	//err = story.Items[0].Comments.Add("xasfdsaf")
	//examples.CheckErr(err)

	//if !examples.UsingSession {
	//	err = instaLogin.Logout()
	//	examples.CheckErr(err)
	//}
}

func getRecent(){

	act := instaLogin.Activity.Recent()

	for act.Next() {
		fmt.Printf("Stories: %d %d\n", len(act.Stories), act.NextID)
		fmt.Println(act.Stories)
		fmt.Println(act.NextID)
	}
	fmt.Println(act.Error())
}






