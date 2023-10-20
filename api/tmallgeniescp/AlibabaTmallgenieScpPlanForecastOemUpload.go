package tmallgeniescp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgeniescp"
)

// Alibabatmallgeniescpplanforecastoemupload 16-供应商预测（OEM-成品）回传接口
// alibaba.tmallgenie.scp.plan.forecast.oem.upload
//
// 供应商预测（OEM-成品）回传接口
func Alibabatmallgeniescpplanforecastoemupload(clt *core.SDKClient, req *tmallgeniescp.AlibabatmallgeniescpplanforecastoemuploadAPIRequest, session string) (*tmallgeniescp.AlibabatmallgeniescpplanforecastoemuploadAPIResponse, error) {
	var resp tmallgeniescp.AlibabatmallgeniescpplanforecastoemuploadAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
