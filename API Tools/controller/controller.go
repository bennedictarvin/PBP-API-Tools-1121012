package controller

import (
	"log"
	"time"

	"context"

	"github.com/go-gomail/gomail"
	"github.com/jasonlvhit/gocron"
)

// GoCron
func RunScheduler() {
	schedule := gocron.NewScheduler(time.UTC)

	schedule.Every(1).Day().Do(func() {
		//Go Routines
		getTodayNews()
		//Mock Redis Example
		//go sendAdvertisement()
	})

	schedule.StartBlocking()
	//time.Sleep(10 * time.Second)
}

// GoMail
func sendMail(user User, news Berita) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", "if-21015@students.ithb.ac.id")
	msg.SetHeader("To", user.Email)
	msg.SetHeader("Subject", news.Judul)
	msg.SetBody("text/html", "<p>"+news.Isi+"</p>")

	n := gomail.NewDialer("smtp.gmail.com", 587, "if-21015@students.ithb.ac.id", "ABC_123456")

	// Send the email
	if err := n.DialAndSend(msg); err != nil {
		log.Println(err)
		return
	}
}

func getTodayNews() {
	db := Connect()
	defer db.Close()

	today := time.Now()

	query := `
			SELECT * FROM berita 
			WHERE tanggal = ?`

	rows, err := db.Query(query, today.Format("2006-01-02"))
	if err != nil {
		log.Println(err)
		return
	}

	var berita Berita
	for rows.Next() {
		if err := rows.Scan(&berita.ID, &berita.Tanggal, &berita.Judul, &berita.Isi); err != nil {
			log.Println(err)
			return
		} else {

			query2 := `SELECT * FROM users`

			rows2, err := db.Query(query2)
			if err != nil {
				log.Println(err)
				return
			}

			var user User
			for rows2.Next() {
				if err := rows2.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
					log.Println(err)
					return
				} else {
					sendMail(user, berita)
				}
			}
		}
	}
}

// Go Redis
func sendAdvertisement() {
	var ctx = context.Background()
	rdb := ConnectRedis()
	db := Connect()
	defer db.Close()

	promoCode, err := rdb.Get(ctx, "promoCode").Result()
	if err != nil {
		panic(err)
	}

	query := `SELECT * FROM users`

	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return
	}
	today := time.Now()
	berita := Berita{ID: 0, Tanggal: today.Format("2006-01-02"), Judul: "Todays Promo Code!", Isi: promoCode}
	var user User
	for rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email); err != nil {
			log.Println(err)
			return
		} else {
			sendMail(user, berita)
		}
	}
}
