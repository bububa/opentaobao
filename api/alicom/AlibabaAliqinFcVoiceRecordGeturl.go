package alicom

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alicom"
)

// Alibabaaliqinfcvoicerecordgeturl 录音文件下载
// alibaba.aliqin.fc.voice.record.geturl
//
// 完成录音文件的下载地址获取操作
func Alibabaaliqinfcvoicerecordgeturl(clt *core.SDKClient, req *alicom.AlibabaaliqinfcvoicerecordgeturlAPIRequest, session string) (*alicom.AlibabaaliqinfcvoicerecordgeturlAPIResponse, error) {
	var resp alicom.AlibabaaliqinfcvoicerecordgeturlAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
