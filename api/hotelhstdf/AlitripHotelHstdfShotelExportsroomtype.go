package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// Alitriphotelhstdfshotelexportsroomtype 导出一个卖家房型下的所有标准房型
// alitrip.hotel.hstdf.shotel.exportsroomtype
//
// 导出一个卖家酒店下的所有标准房型
func Alitriphotelhstdfshotelexportsroomtype(clt *core.SDKClient, req *hotelhstdf.AlitriphotelhstdfshotelexportsroomtypeAPIRequest, session string) (*hotelhstdf.AlitriphotelhstdfshotelexportsroomtypeAPIResponse, error) {
	var resp hotelhstdf.AlitriphotelhstdfshotelexportsroomtypeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
