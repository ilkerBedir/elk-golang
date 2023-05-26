package main

import (
  "os"
  log "github.com/sirupsen/logrus"
  "net/http"
  "github.com/go-chi/chi"
	"github.com/go-chi/cors"
)
func main() {
  log.SetFormatter(&log.JSONFormatter{
    FieldMap: log.FieldMap{                               
      log.FieldKeyTime:  "@timestamp",            
      log.FieldKeyMsg:   "message",
    },
  })
  log.SetLevel(log.TraceLevel)

  file, err := os.OpenFile("./logs/go.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
  if err == nil {
    log.SetOutput(file)
  }
  defer file.Close()
  hostName,_:=os.Hostname()
  fields := log.Fields{"computername": hostName}
  logger:=log.WithFields(fields)
  logger.Info("Starting Demo Program Server...")
  portString:="8080"
  router:=chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET","POST","PUT","DELETE","OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))
  router.Get("/hello",func(w http.ResponseWriter,r *http.Request){
      log.Debug("Logger says hello")
  })
  router.Get("/olleh",func(w http.ResponseWriter,r *http.Request){
    log.Debug("Logger says olleh")
})
router.Get("/1",func(w http.ResponseWriter,r *http.Request){
  log.Debug("Logger says one")
})
	srv:=&http.Server{
		Handler:router,
		Addr: ":"+portString,
	}
	log.Info("Server starting on port "+portString)
	err = srv.ListenAndServe()
	if err != nil {
		log.Error(err)
	}
}