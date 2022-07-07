package api

import (
	"encoding/json"
	"go-code/log"
	"go-code/utility"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
}

// Response ...
type Response struct {
	Message  string      `json:"message" example:"in depth explanation"`
	Response interface{} `json:"response" example:"final ans"`
}

type ReqJson struct {
	Array      []int `json:"array"`
	Quantifier int   `json:"quantifier"`
}

type VersionReq struct {
	Version1 string `json:"version1"`
	Version2 string `json:"version2"`
}

func (ctrl Controller) Min(c *gin.Context) {
	inputReq, err := c.GetRawData()
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid payload"})
		return
	}
	req := ReqJson{}
	err = json.Unmarshal(inputReq, &req)
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid request"})
		return
	}
	if len(req.Array) < req.Quantifier {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "quantifier cannot be greater than length of numbers"})
		return
	}
	resp := utility.MinMaxNumbers(req.Array, req.Quantifier, true)
	finalResponse := Response{
		Message:  "Min Numbers",
		Response: resp,
	}
	log.Info("Successful Response", log.Fields{
		"minNumber": resp,
	})
	c.JSON(http.StatusOK, finalResponse)
}

func (ctrl Controller) Max(c *gin.Context) {
	inputReq, err := c.GetRawData()
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid payload"})
		return
	}
	req := ReqJson{}
	err = json.Unmarshal(inputReq, &req)
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid request"})
		return
	}
	if len(req.Array) < req.Quantifier {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "quantifier cannot be greater than length of numbers"})
		return
	}
	resp := utility.MinMaxNumbers(req.Array, req.Quantifier, false)
	finalResponse := Response{
		Message:  "Max Numbers",
		Response: resp,
	}
	log.Info("Successful Response", log.Fields{
		"maxNumber": resp,
	})
	c.JSON(http.StatusOK, finalResponse)
}

func (ctrl Controller) Median(c *gin.Context) {
	inputReq, err := c.GetRawData()
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid payload"})
		return
	}
	req := ReqJson{}
	err = json.Unmarshal(inputReq, &req)
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid request"})
		return
	}
	resp := utility.Median(req.Array)
	finalResponse := Response{
		Message:  "Median",
		Response: resp,
	}
	log.Info("Successful Response", log.Fields{
		"median": resp,
	})
	c.JSON(http.StatusOK, finalResponse)
}

func (ctrl Controller) Average(c *gin.Context) {
	inputReq, err := c.GetRawData()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid payload"})
		return
	}
	req := ReqJson{}
	err = json.Unmarshal(inputReq, &req)
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid request"})
		return
	}
	resp := utility.Average(req.Array)
	finalResponse := Response{
		Message:  "Average",
		Response: resp,
	}
	log.Info("Successful Response", log.Fields{
		"average": resp,
	})
	c.JSON(http.StatusOK, finalResponse)
}

func (ctrl Controller) Percentaile(c *gin.Context) {
	inputReq, err := c.GetRawData()
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid payload"})
		return
	}
	req := ReqJson{}
	err = json.Unmarshal(inputReq, &req)
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid request"})
		return
	}
	resp := utility.Percentaile(req.Array, req.Quantifier)
	finalResponse := Response{
		Message:  "Percentaile",
		Response: resp,
	}
	log.Info("Successful Response", log.Fields{
		"percentaile": resp,
	})
	c.JSON(http.StatusOK, finalResponse)
}

func (ctrl Controller) Version(c *gin.Context) {
	inputReq, err := c.GetRawData()
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid payload"})
		return
	}
	req := VersionReq{}
	err = json.Unmarshal(inputReq, &req)
	if err != nil {
		log.Error(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "invalid request"})
		return
	}
	if req.Version1 == "" || req.Version2 == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, Response{Message: "empty version numbers"})
		return
	}
	resp := utility.CompareVersion(req.Version1, req.Version2)
	finalResponse := Response{
		Message:  "Version Check",
		Response: resp,
	}
	log.Info("Successful Response", log.Fields{
		"version check": resp,
	})
	c.JSON(http.StatusOK, finalResponse)
}
