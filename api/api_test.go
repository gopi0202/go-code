package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var (
	res *httptest.ResponseRecorder
	con *gin.Context
	eng *gin.Engine
)

func setUp() {
	gin.SetMode(gin.TestMode)
	res = httptest.NewRecorder()
	con, eng = gin.CreateTestContext(res)
	Routes(eng.Group("/app"))
}

func TestMin(t *testing.T) {
	setUp()
	cases := []struct {
		name   string
		req    ReqJson
		output Response
	}{
		{
			name: "Success case one quantifier",
			req: ReqJson{
				Array:      []int{1, 2, 3},
				Quantifier: 1,
			},
			output: Response{
				Response: []int{1},
				Message:  "Min Numbers",
			},
		},
		{
			name: "Success case multiple quantifier",
			req: ReqJson{
				Array:      []int{1, 2, 3},
				Quantifier: 2,
			},
			output: Response{
				Response: []int{1, 2},
				Message:  "Min Numbers",
			},
		},
		{
			name: "Quantifier greater than array length",
			req: ReqJson{
				Array:      []int{1, 2, 3},
				Quantifier: 5,
			},
			output: Response{
				Response: "",
				Message:  "quantifier cannot be greater than length of numbers",
			},
		},
		{
			name: "Big Numbers",
			req: ReqJson{
				Array:      []int{1000, 2222, 4444, 222},
				Quantifier: 2,
			},
			output: Response{
				Response: []int{222, 1000},
				Message:  "Min numbers",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, _ := json.Marshal(c.req)
			req, _ := http.NewRequest("GET", "/min", bytes.NewBufferString(string(data)))
			req.Header.Set("Content-Type", "application/json")
			eng.ServeHTTP(res, req)
			var response Response
			err := json.Unmarshal([]byte(res.Body.Bytes()), &response)
			if err != nil {
				return
			}
			assert.Equal(t, response, c.output)
		})
	}

}

func TestMax(t *testing.T) {
	setUp()
	cases := []struct {
		name   string
		req    ReqJson
		output Response
	}{
		{
			name: "Success case one quantifier",
			req: ReqJson{
				Array:      []int{1, 2, 3},
				Quantifier: 1,
			},
			output: Response{
				Response: []int{3},
				Message:  "Max Numbers",
			},
		},
		{
			name: "Success case multiple quantifier",
			req: ReqJson{
				Array:      []int{1, 2, 3},
				Quantifier: 2,
			},
			output: Response{
				Response: []int{2, 3},
				Message:  "Max Numbers",
			},
		},
		{
			name: "Quantifier greater than array length",
			req: ReqJson{
				Array:      []int{1, 2, 3},
				Quantifier: 5,
			},
			output: Response{
				Response: []int{1, 2},
				Message:  "quantifier cannot be greater than length of numbers",
			},
		},
		{
			name: "Big Numbers",
			req: ReqJson{
				Array:      []int{1000, 2222, 4444, 222},
				Quantifier: 2,
			},
			output: Response{
				Response: []int{4444, 2222},
				Message:  "Max numbers",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, _ := json.Marshal(c.req)
			req, _ := http.NewRequest("GET", "/max", bytes.NewBufferString(string(data)))
			req.Header.Set("Content-Type", "application/json")
			eng.ServeHTTP(res, req)
			var response Response
			err := json.Unmarshal([]byte(res.Body.Bytes()), &response)
			if err != nil {
				return
			}
			assert.Equal(t, response, c.output)
		})
	}

}

func TestAverage(t *testing.T) {
	setUp()
	cases := []struct {
		name   string
		req    ReqJson
		output Response
	}{
		{
			name: "Average",
			req: ReqJson{
				Array: []int{1, 2, 3},
			},
			output: Response{
				Response: 1,
				Message:  "Average",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, _ := json.Marshal(c.req)
			req, _ := http.NewRequest("GET", "/avg", bytes.NewBufferString(string(data)))
			req.Header.Set("Content-Type", "application/json")
			eng.ServeHTTP(res, req)
			var response Response
			err := json.Unmarshal([]byte(res.Body.Bytes()), &response)
			if err != nil {
				return
			}
			assert.Equal(t, response, c.output)
		})
	}

}

func TestMedian(t *testing.T) {
	setUp()
	cases := []struct {
		name   string
		req    ReqJson
		output Response
	}{
		{
			name: "Average",
			req: ReqJson{
				Array: []int{1, 2, 3, 4, 5},
			},
			output: Response{
				Response: 3,
				Message:  "Median",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, _ := json.Marshal(c.req)
			req, _ := http.NewRequest("GET", "/median", bytes.NewBufferString(string(data)))
			req.Header.Set("Content-Type", "application/json")
			eng.ServeHTTP(res, req)
			var response Response
			err := json.Unmarshal([]byte(res.Body.Bytes()), &response)
			if err != nil {
				return
			}
			assert.Equal(t, response, c.output)
		})
	}

}

func TestPercentaile(t *testing.T) {
	setUp()
	cases := []struct {
		name   string
		req    ReqJson
		output Response
	}{
		{
			name: "Success case one quantifier",
			req: ReqJson{
				Array:      []int{1, 2, 3},
				Quantifier: 1,
			},
			output: Response{
				Response: []int{3},
				Message:  "Percentaile",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, _ := json.Marshal(c.req)
			req, _ := http.NewRequest("GET", "/percentaile", bytes.NewBufferString(string(data)))
			req.Header.Set("Content-Type", "application/json")
			eng.ServeHTTP(res, req)
			var response Response
			err := json.Unmarshal([]byte(res.Body.Bytes()), &response)
			if err != nil {
				return
			}
			assert.Equal(t, response, c.output)
		})
	}

}

func TestVersion(t *testing.T) {
	setUp()
	cases := []struct {
		name   string
		req    VersionReq
		output Response
	}{
		{
			name: "Success case one quantifier",
			req: VersionReq{
				Version1: "2.1.1",
				Version2: "1.1.3",
			},
			output: Response{
				Response: 1,
				Message:  "Version Check",
			},
		},
		{
			name: "Success case multiple quantifier",
			req: VersionReq{
				Version1: "12.1.1",
				Version2: "12.1.3",
			},
			output: Response{
				Response: -1,
				Message:  "Version Check",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			data, _ := json.Marshal(c.req)
			req, _ := http.NewRequest("GET", "/version", bytes.NewBufferString(string(data)))
			req.Header.Set("Content-Type", "application/json")
			eng.ServeHTTP(res, req)
			var response Response
			err := json.Unmarshal([]byte(res.Body.Bytes()), &response)
			if err != nil {
				return
			}
			assert.Equal(t, response, c.output)
		})
	}

}
