package main

import (
	"GinDemo/src/pkg/setting"
	"GinDemo/src/routers"
	v1 "GinDemo/src/routers/api/v1"
	"fmt"
	"net/http"
)



func main() {
	router := routers.InitRouter()
	apiv1:=router.Group("api/v1")
	{
		apiv1.GET("/tags",v1.GetTags)
		apiv1.POST("/tags",v1.AddTags)
		apiv1.PUT("/tags/:id",v1.EditTags)
		apiv1.DELETE("/tags/:id",v1.DeleteTags)
	}


	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()

}
