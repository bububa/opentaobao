package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// Alitriphotelhstdfshotelexnotmatchroom 导出一个hid下所有未匹配rid的接口
// alitrip.hotel.hstdf.shotel.exnotmatchroom
//
// 导出一个卖家hid下所有未匹配的rid信息，包括rid，名称、英文名称、床型、窗型、面积、对外系统id
func Alitriphotelhstdfshotelexnotmatchroom(clt *core.SDKClient, req *hotelhstdf.AlitriphotelhstdfshotelexnotmatchroomAPIRequest, session string) (*hotelhstdf.AlitriphotelhstdfshotelexnotmatchroomAPIResponse, error) {
	var resp hotelhstdf.AlitriphotelhstdfshotelexnotmatchroomAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
