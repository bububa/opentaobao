package media

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/media"
)

// TaobaovasservicegetServTimes 查询某个用户图片空间的使用情况
// taobao.vas.service.getServTimes
//
// 查询某个用户图片空间的使用情况
func TaobaovasservicegetServTimes(clt *core.SDKClient, req *media.TaobaovasservicegetServTimesAPIRequest, session string) (*media.TaobaovasservicegetServTimesAPIResponse, error) {
	var resp media.TaobaovasservicegetServTimesAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
