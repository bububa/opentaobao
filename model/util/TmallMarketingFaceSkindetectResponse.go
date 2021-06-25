package util

import (
    "github.com/bububa/opentaobao/model"
)

/* 
肌肤检测 APIResponse
tmall.marketing.face.skindetect

提供人脸肌肤属性报告
*/
type TmallMarketingFaceSkindetectAPIResponse struct {
    model.CommonResponse
    Response *TmallMarketingFaceSkindetectResponse `json:"tmall_marketing_face_skindetect_response,omitempty"`
}

type TmallMarketingFaceSkindetectResponse struct {

    // { 	"code": "top api 状态码 200表示成功 ", 	"msg": "top api 返回描述说明 ", 	"data": { 		"detect_time": "检测时间戳 ", 		"gender": "0为男性1为女性 ", 		"age": "年龄 ", 		"color_level": "肤色等级，0~7，数字越小越白 ", 		"hue_level": " 色调等级，0~2，冷色，中性，暖色", 		"oil_cheeck": "脸颊含油 ", 		"oil_t_area": "T区含油 ", 		"oil_chin": "下巴含油 ", 		"oil_level": "总体出油情况, 0-干性，1-偏干，2-中性，3-混合油性，4-偏油，5-油性 ", 		"smooth_level": " 光滑度等级，0~3，数字越小越光滑 ", 		"acne_level": "痘痘严重程度，0~3，0表示没有，1~3表示轻度、中度、重度 ", 		"pore_level": " 毛孔粗细等级，0~2 - 细致，较粗大，粗大 ", 		"blackheads": "黑头数量 ", 		"black_level": "黑头严重程度，0~3，0表示没有，1~3表示轻度、中度、重度 ", 		"acne_loc": "痘痘坐标x,y，半径大小r(x,y,r)", 		"score": "颜值分", 		"r_face": "人脸区域坐标（x,y,w,h）", 		"code": "状态码,0表示成功" 	} }
    DetectResult   string `json:"detect_result,omitempty"`

}
