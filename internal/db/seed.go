package db

import (
	"context"
	"fmt"
	"log"
	"math/rand/v2"

	"github.com/codepnw/social/internal/store"
)

var usernames = []string{
	"UserAlpha01", "BetaGamer22", "GammaCoder33", "DeltaRider44", "EpsilonHero55",
	"ZetaHunter66", "EtaWizard77", "ThetaSamurai88", "IotaNinja99", "KappaPilot11",
	"LambdaRogue12", "MuWarrior13", "NuKnight14", "XiSage15", "OmicronHero16",
	"PiExplorer17", "RhoSeeker18", "SigmaDiver19", "TauDreamer20", "UpsilonLegend21",
	"PhiRacer22", "ChiCreator23", "PsiThinker24", "OmegaCrafter25", "StormBreaker26",
	"SkyWatcher27", "StarVoyager28", "MoonJumper29", "GalaxyRuler30", "NovaWanderer31",
	"CometChaser32", "NebulaRunner33", "AstroGuru34", "OrbitSailor35", "CosmoAdventurer36",
	"BlackHoleSurfer37", "TimeTraveler38", "QuantumHacker39", "PhotonSlasher40",
	"FusionStriker41", "GravityCrusher42", "PlasmaWalker43", "DarkMatterWizard44",
	"SolarChampion45", "RocketRider46", "AsteroidHunter47", "EclipseShifter48",
	"StellarMage49", "InfinityRanger50",
}

var titles = []string{
	"Go for Beginners", "Mastering APIs", "Concurrency in Go",
	"Web Dev Basics", "Understanding Interfaces", "Docker and Go",
	"Unit Testing Tips", "Building CLI Tools", "Clean Code Practices",
	"Effective Error Handling", "Goroutines Explained", "JSON in Go",
	"Working with Databases", "HTTP in Go", "Using Middleware",
	"File Handling Basics", "Go Dependency Management", "RESTful APIs",
	"Debugging Go Apps", "Scaling Go Applications",
}

var contents = []string{
	"Discover the secrets of Go programming.",
	"Learn how to scale your applications effortlessly.",
	"Concurrency made simple with goroutines.",
	"The ultimate guide to error handling in Go.",
	"Understanding Go's type system.",
	"Exploring web development with Go.",
	"Mastering RESTful APIs in Go.",
	"Clean architecture for maintainable code.",
	"Debugging techniques for Go developers.",
	"Introduction to Go's garbage collector.",
	"Best practices for Go unit testing.",
	"Optimizing performance in Go applications.",
	"Working with files and directories in Go.",
	"Understanding Go's `sync` package.",
	"Building CLI tools using Go.",
	"Managing dependencies with `go mod`.",
	"Practical tips for Go beginners.",
	"Structuring large Go projects.",
	"Using channels effectively in Go.",
	"Exploring middleware in web frameworks.",
}

var tags = []string{
	"go", "programming", "webdev", "api", "database",
	"concurrency", "testing", "docker", "clean-code", "restful-api",
	"scalability", "debugging", "json", "middleware", "cli-tools",
	"file-handling", "performance", "architecture", "goroutines", "http",
	"devops", "microservices", "security", "deployment", "open-source",
	"golang-tips", "coding", "backend", "tech-blog", "web-architecture",
}

var comments = []string{
	"Great work!",
	"Needs some improvements.",
	"Interesting approach!",
	"Consider refactoring this part.",
	"Awesome performance!",
	"Good use of patterns.",
	"Code could use more comments.",
	"Clear and concise!",
	"Logic is a bit unclear here.",
	"Nice optimization!",
	"Consider edge cases.",
	"Well done overall!",
	"Can be improved for readability.",
	"Excellent documentation.",
	"Check for potential errors.",
	"Good variable naming.",
	"Try modularizing the functions.",
	"Nice use of Go idioms.",
	"Great unit tests!",
	"Focus on performance here.",
}

func Seed(store store.Storage) {
	ctx := context.Background()

	users := generateUsers(100)
	for _, user := range users {
		if err := store.Users.Create(ctx, user); err != nil {
			log.Println("error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("error creating post:", err)
			return
		}
	}

	comments := generateComments(500, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("error creating comment:", err)
			return
		}
	}

	log.Println("seeding completed")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
			Password: "123123",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)
	for i := 0; i < num; i++ {
		user := users[rand.IntN(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.IntN(len(titles))],
			Content: contents[rand.IntN(len(contents))],
			Tags:    []string{
				tags[rand.IntN(len(tags))],
				tags[rand.IntN(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	cms := make([]*store.Comment, num)
	for i := 0; i < num; i++ {
		cms[i] = &store.Comment{
			PostID: posts[rand.IntN(len(posts))].ID,
			UserID: users[rand.IntN(len(users))].ID,
			Content: comments[rand.IntN(len(comments))],
		}
	}
	return cms
}
