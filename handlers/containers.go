package handlers

import (
	"fmt"
	"net/http"
	"os"

	"proxmox-api/config"
	"proxmox-api/models"
	"proxmox-api/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	templateDir = "/terraform/template"
	workspace   = "/workspace/containers"
)

func CreateContainer(c *gin.Context) {

	var req models.CreateContainerRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.New().String()

	workDir := fmt.Sprintf("%s/%s", workspace, id)

	os.MkdirAll(workDir, os.ModePerm)

	err := services.CopyDir(templateDir, workDir)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	tfvars := fmt.Sprintf(`
container_name="%s"
cpu=%d
memory=%d
`, req.Name, req.CPU, req.Memory)

	os.WriteFile(workDir+"/terraform.tfvars", []byte(tfvars), 0644)

	err = services.TerraformInit(workDir, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = services.TerraformApply(workDir)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	container := models.Container{
		UUID:   id,
		Name:   req.Name,
		CPU:    req.CPU,
		Memory: req.Memory,
		Status: "running",
	}

	config.DB.Create(&container)

	c.JSON(200, container)
}

func DeleteContainer(c *gin.Context) {

	id := c.Param("uuid")

	var container models.Container

	if err := config.DB.Where("uuid = ?", id).First(&container).Error; err != nil {
		c.JSON(404, gin.H{"error": "container not found"})
		return
	}

	workDir := fmt.Sprintf("%s/%s", workspace, id)

	err := services.TerraformInit(workDir, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	err = services.TerraformDestroy(workDir)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	os.RemoveAll(workDir)

	container.Status = "destroyed"
	config.DB.Save(&container)

	c.JSON(200, gin.H{
		"uuid":   id,
		"status": "destroyed",
	})
}

func ListContainers(c *gin.Context) {

	var containers []models.Container

	config.DB.Find(&containers)

	c.JSON(200, containers)
}