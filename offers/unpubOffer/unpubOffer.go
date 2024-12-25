package unpuboffer

import (
	"be/conf"
	offerstr "be/offers/offerStr"
	userstr "be/userStr"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UnPubOffer(c *gin.Context) {
	id, exists := c.Get("id")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Id not found in context"})
		return
	}
	id = int(id.(float64))

	var user userstr.User
	user.Id = id.(int)

	var offer offerstr.Offer
	if err := c.ShouldBindJSON(&offer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	errOffer := conf.DB.Where("id = ?", offer.Id).First(&offer).Error
	if errOffer != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "offer not found"})
		log.Println(http.StatusNotFound, gin.H{"error": "offer not found"})
		return
	}

	log.Println("offer.Id_c ", offer.Id_c)
	log.Println("user.Id ", user.Id)

	if offer.Id_c != user.Id {
		c.JSON(http.StatusNotFound, gin.H{"error": "offers not found"})
		log.Println(http.StatusNotFound, gin.H{"error": "Айди пользователя не совпадает с айди_с офера"})
		return
	}

	offer.IsPub = false

	err := conf.DB.Save(&offer).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "offers not save"})
		log.Println(http.StatusInternalServerError, gin.H{"error": "offers not save"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Offer unpub"})
}
