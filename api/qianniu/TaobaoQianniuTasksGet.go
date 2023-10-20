package qianniu

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qianniu"
)

// Taobaoqianniutasksget 获取指定的任务
// taobao.qianniu.tasks.get
//
// 获取指定的任务，可用的参数组合：&lt;br/&gt;task_ids + need_meta + fields：精确查找&lt;br/&gt;biz_type + sub_biz_type + biz_ids + need_meta + fields：按照业务ID查找&lt;br/&gt;biz_type + sub_biz_type + sender_uid + need_meta + fields：按照发起者查找&lt;br/&gt;biz_type + sub_biz_type + receiver_uid + need_meta + fields：按照执行者查找&lt;br/&gt;biz_type+modify_start_time+modify_end_time+fields:能支持指定修改时间的查询，用于增量查询等
func Taobaoqianniutasksget(clt *core.SDKClient, req *qianniu.TaobaoqianniutasksgetAPIRequest, session string) (*qianniu.TaobaoqianniutasksgetAPIResponse, error) {
	var resp qianniu.TaobaoqianniutasksgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
