package perfect

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/perfect"
)

// Alibabaperfectperformanceitemquery 商品完美履约信息查询
// alibaba.perfect.performance.item.query
//
// 同城零售商品完美履约信息查询
func Alibabaperfectperformanceitemquery(clt *core.SDKClient, req *perfect.AlibabaperfectperformanceitemqueryAPIRequest, session string) (*perfect.AlibabaperfectperformanceitemqueryAPIResponse, error) {
	var resp perfect.AlibabaperfectperformanceitemqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
