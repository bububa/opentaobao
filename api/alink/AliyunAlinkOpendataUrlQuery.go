package alink

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alink"
)

// Aliyunalinkopendataurlquery 开放数据授权访问URL查询
// aliyun.alink.opendata.url.query
//
// 厂商数据授权访问URL查询
func Aliyunalinkopendataurlquery(clt *core.SDKClient, req *alink.AliyunalinkopendataurlqueryAPIRequest, session string) (*alink.AliyunalinkopendataurlqueryAPIResponse, error) {
	var resp alink.AliyunalinkopendataurlqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
