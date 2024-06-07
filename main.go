package main

import (
	"encoding/xml"
	"flag"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type IPResponse struct {
	XMLName xml.Name `xml:"response" json:"-"`
	IP      string   `xml:"ip" json:"ip"`
}

func main() {
	// Define a CLI argument for the template directory
	templateDir := flag.String("templateDir", "./templates", "directory path for HTML templates")
	staticDir := flag.String("staticDir", "./static", "directory path for static files")
	flag.Parse()

	r := gin.Default()

	// Serve static files from the "static" directory
	r.Static("/static", *staticDir)

	// HTML template route with the specified or default directory
	r.LoadHTMLGlob(*templateDir + "/*")
	r.GET("/", func(c *gin.Context) {
		ip := getClientIP(c.Request)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"ip": ip,
		})
	})

	// Plain text route
	r.GET("/ip/text", func(c *gin.Context) {
		ip := getClientIP(c.Request)
		c.String(http.StatusOK, ip)
	})

	// JSON route
	r.GET("/ip/json", func(c *gin.Context) {
		ip := getClientIP(c.Request)
		c.JSON(http.StatusOK, IPResponse{IP: ip})
	})

	// XML route
	r.GET("/ip/xml", func(c *gin.Context) {
		ip := getClientIP(c.Request)
		c.XML(http.StatusOK, IPResponse{IP: ip})
	})

	r.Run(":8080")
}

// getClientIP determines the real client IP, even behind proxies like Cloudflare and Nginx.
func getClientIP(req *http.Request) string {
	// Check Cloudflare's CF-Connecting-IP header
	if ip := req.Header.Get("CF-Connecting-IP"); ip != "" {
		return ip
	}

	// Check headers that might contain proxy IP information
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		// X-Forwarded-For can contain multiple IPs, the first one is the client IP
		return strings.Split(ip, ",")[0]
	}

	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}

	// Fallback to the request's RemoteAddr
	ip := strings.Split(req.RemoteAddr, ":")[0]
	return ip
}
