package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/zoellert/cloudflare-dyndns/services"
	"net/http"
)

func UpdateDNSRecord(c *gin.Context) {
	hostname := c.Query("hostname")
	ip := c.Query("myip")

	if len(hostname) < 1 || len(ip) < 1 {
		c.String(http.StatusBadRequest, "Parameter hostname and/or ip missing")
		return
	}

	err, hasChanged := services.UpdateDNSRecord(hostname, ip)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to update dns record")
		return
	}

	if hasChanged {
		c.String(http.StatusOK, "good "+ip)
	} else {
		c.String(http.StatusOK, "nochg "+ip)
	}
}
