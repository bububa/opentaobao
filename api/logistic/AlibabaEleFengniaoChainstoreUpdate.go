package logistic

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/logistic"
)

// Alibabaelefengniaochainstoreupdate 修改门店信息接口
// alibaba.ele.fengniao.chainstore.update
//
// 修改门店的经纬度，文本地址，电话，门店名
func Alibabaelefengniaochainstoreupdate(clt *core.SDKClient, req *logistic.AlibabaelefengniaochainstoreupdateAPIRequest, session string) (*logistic.AlibabaelefengniaochainstoreupdateAPIResponse, error) {
	var resp logistic.AlibabaelefengniaochainstoreupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
