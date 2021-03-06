package process

import (
	"Vua-Crawler/spider/model"
	"net/http"
)

// 处理返回的页面接口
type Process interface {
	Process(response *http.Response) (*model.Page, error)
}
