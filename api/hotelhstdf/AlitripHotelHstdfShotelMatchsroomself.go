package hotelhstdf

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/hotelhstdf"
)

// Alitriphotelhstdfshotelmatchsroomself 匹配标准房型以及卖家房型
// alitrip.hotel.hstdf.shotel.matchsroomself
//
// 匹配卖家房型以及标准房型，返回匹配结果
func Alitriphotelhstdfshotelmatchsroomself(clt *core.SDKClient, req *hotelhstdf.AlitriphotelhstdfshotelmatchsroomselfAPIRequest, session string) (*hotelhstdf.AlitriphotelhstdfshotelmatchsroomselfAPIResponse, error) {
	var resp hotelhstdf.AlitriphotelhstdfshotelmatchsroomselfAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
