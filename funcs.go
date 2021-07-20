package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func handleDownload(c *gin.Context) {
	cfg_name := c.Param("cfg_name")

	if is_updated, ok := config_version_changed[cfg_name]; !ok {
		config_version_changed[cfg_name] = true
	} else if !is_updated {
		c.JSON(http.StatusOK,
			gin.H{
				"msg": "Not updated",
			})
		return
	}

	if _, err := os.Open("./cfgs/" + cfg_name); err != nil {
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"msg": "Download failed",
			})
		return
	}

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", cfg_name))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./cfgs/" + cfg_name)

}

func handleUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		log.Println("ERROR: upload file failed. ", err)
		c.JSON(http.StatusInternalServerError,
			gin.H{
				"msg": fmt.Sprintf("ERROR: upload file failed. %s", err),
			})
		return
	}
	dst := fmt.Sprintf("./cfgs/" + file.Filename)
	err = c.SaveUploadedFile(file, dst)

	config_version_changed[file.Filename] = true
	c.JSON(http.StatusOK, gin.H{
		"msg":      "succ",
		"filepath": dst,
	})
}
