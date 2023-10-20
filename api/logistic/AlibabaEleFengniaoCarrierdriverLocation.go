package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaelefengniaocarrierdriverlocation 查询骑手当前位置
// alibaba.ele.fengniao.carrierdriver.location
//
// 查询骑手当前位置
func Alibabaelefengniaocarrierdriverlocation(clt *core.SDKClient, req *logistic.AlibabaelefengniaocarrierdriverlocationAPIRequest, session string) (*logistic.AlibabaelefengniaocarrierdriverlocationAPIResponse, error) {
	var resp logistic.AlibabaelefengniaocarrierdriverlocationAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
