package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxidriverlocationget 司机位置
// alibaba.happytrip.taxi.driver.location.get
//
// 获取司机实时位置
func Alibabahappytriptaxidriverlocationget(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxidriverlocationgetAPIRequest, session string) (*happytrip.AlibabahappytriptaxidriverlocationgetAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxidriverlocationgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
