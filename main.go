/* ----------------------------------
*  @author suyame 2022-07-06 9:44:00
*  Crazy for Golang !!!
*  IDE: GoLand
*-----------------------------------*/
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/mysql/new", func(c *gin.Context) {
		//
		name := c.Query("name")
		user, err := AddUserInfo(name, "123456")
		if err == nil {
			addCache(&user)
			c.JSON(200, gin.H{
				"Message": "successfully!",
			})
		} else {
			c.JSON(500, gin.H{
				"Message": "failed!",
			})
		}
	})

	r.GET("/mysql/get", func(c *gin.Context) {
		//
		name := c.Query("name")
		user, ok := FindUserInfo(name)
		if ok {
			c.String(200, "%#v", user)
		} else {
			c.String(404, "%v not found! ", name)
		}
	})

	r.GET("/mysql/update", func(c *gin.Context) {
		//
		name := c.Query("name")
		newname := c.DefaultQuery("newname", "newname")
		user, ok := FindUserInfo(name)
		if !ok {
			c.String(404, "%v not found! ", name)
		} else {
			user.Name = newname
			deleteCache(name)
			addCache(&user)
			UpdateUser(user)
		}
	})

	r.GET("/mysql/delete", func(c *gin.Context) {
		//
		name := c.Query("name")
		err := DeleteUser(name)
		if err != nil {
			c.String(500, "%v delete failed!", err)
		} else {
			// 从redis删除
			deleteCache(name)
			c.String(200, "successfully!")
		}
	})

	r.GET("/redis/get", func(c *gin.Context) {
		name := c.Query("name")
		user, err := getUserFormCache(name)
		if err == nil {
			c.String(200, "%#v", user)
		} else {
			c.String(500, "%v not found! %v", name, err)
		}
	})

	r.Run()
}
