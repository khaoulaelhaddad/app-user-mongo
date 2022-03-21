package rest

import (
	"context"
	"log"
	"time"
	"userApp/rest/handlers"
	"userApp/storage"
	mg "userApp/storage/mg"
	"userApp/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	Engine  *gin.Engine
	Log     *logrus.Logger
	Storage storage.Storage
}

func (s *Server) SetupEngin() {
	s.Engine = gin.Default()
}

func (s *Server) SetupLog() {
	s.Log = logrus.New()
	s.Log.SetLevel(logrus.InfoLevel)
}

func (s *Server) SetupStorage() {

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(utils.ApplyURI).
		SetServerAPIOptions(serverAPIOptions)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	s.Storage = &mg.DbConnexion{Db: client.Database("user_database")}

}

func (s *Server) SetupCors() {
	s.Engine.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "PATCH", "HEAD", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,

		MaxAge: 90 * 24 * time.Hour,
	}))
}

func (s *Server) SetupRoute() {
	baseGroup := s.Engine.Group("")
	handlers.PublicRoutes(baseGroup)

	apiGroup := baseGroup.Group("/api/v1")
	handlers.UserRoutes(apiGroup, s.Storage)
}

func NewServer() (s *Server) {
	s = &Server{}
	s.SetupEngin()
	s.SetupLog()
	s.SetupStorage()
	s.SetupCors()

	s.SetupRoute()
	return s
}

func (s *Server) Run() {
	s.Engine.Run(utils.ServerHost)
}
