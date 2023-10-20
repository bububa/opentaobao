package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaelefengniaoservicepackagequery 预采购服务包查询接口
// alibaba.ele.fengniao.service.package.query
//
// 查询门店所在经纬度可用服务包的接口
func Alibabaelefengniaoservicepackagequery(clt *core.SDKClient, req *logistic.AlibabaelefengniaoservicepackagequeryAPIRequest, session string) (*logistic.AlibabaelefengniaoservicepackagequeryAPIResponse, error) {
	var resp logistic.AlibabaelefengniaoservicepackagequeryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
