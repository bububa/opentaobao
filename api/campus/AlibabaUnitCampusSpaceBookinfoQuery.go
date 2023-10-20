package campus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/campus"
)

// Alibabaunitcampusspacebookinfoquery 环路资源信息查询单元环境
// alibaba.unit.campus.space.bookinfo.query
//
// 环路资源信息查询单元环境
func Alibabaunitcampusspacebookinfoquery(clt *core.SDKClient, req *campus.AlibabaunitcampusspacebookinfoqueryAPIRequest, session string) (*campus.AlibabaunitcampusspacebookinfoqueryAPIResponse, error) {
	var resp campus.AlibabaunitcampusspacebookinfoqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
