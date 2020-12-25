package wpub

import "encoding/xml"

type MsgType string

type EventType string

const (
	MsgTypeText            MsgType = "text"                      //表示文本消息
	MsgTypeImage                   = "image"                     //表示文本消息
	MsgTypeVoice                   = "voice"                     //表示语音消息
	MsgTypeVideo                   = "video"                     //表示视频消息
	MsgTypeMiniprogrampage         = "miniprogrampage"           //表示小程序卡片消息
	MsgTypeShortVideo              = "shortvideo"                //短视频消息[限接收]
	MsgTypeLocation                = "location"                  //坐标消息[限接收]
	MsgTypeLink                    = "link"                      //链接消息[限接收]
	MsgTypeMusic                   = "music"                     //音乐消息[限回复]
	MsgTypeNews                    = "news"                      //图文消息[限回复]
	MsgTypeTransfer                = "transfer_customer_service" //消息转发到客服
	MsgTypeEvent                   = "event"                     //事件推送消息
)

const (
	EventTypeSubscribe             EventType = "subscribe"             //订阅
	EventTypeUnsubscribe                     = "unsubscribe"           //取消订阅
	EventTypeScan                            = "SCAN"                  //用户已经关注公众号，则微信会将带场景值扫描事件推送给开发者
	EventTypeLocation                        = "LOCATION"              //上报地理位置事件
	EventTypeClick                           = "CLICK"                 //点击菜单拉取消息时的事件推送
	EventTypeView                            = "VIEW"                  //点击菜单跳转链接时的事件推送
	EventTypeScancodePush                    = "scancode_push"         //扫码推事件的事件推送
	EventTypeScancodeWaitmsg                 = "scancode_waitmsg"      //扫码推事件且弹出“消息接收中”提示框的事件推送
	EventTypePicSysphoto                     = "pic_sysphoto"          //弹出系统拍照发图的事件推送
	EventTypePicPhotoOrAlbum                 = "pic_photo_or_album"    //弹出拍照或者相册发图的事件推送
	EventTypePicWeixin                       = "pic_weixin"            //弹出微信相册发图器的事件推送
	EventTypeLocationSelect                  = "location_select"       // 弹出地理位置选择器的事件推送
	EventTypeTemplateSendJobFinish           = "TEMPLATESENDJOBFINISH" //发送模板消息推送通知
	EventTypeWxaMediaCheck                   = "wxa_media_check"       //异步校验图片/音频是否含有违法违规内容推送事件
)

type Event struct {
	XMLName      xml.Name  `xml:"xml"`
	ToUserName   string    `xml:"ToUserName"`   //开发者微信号
	FromUserName string    `xml:"FromUserName"` //发送方微信号，若为普通用户，则是一个OpenID
	CreateTime   int64     `xml:"CreateTime"`   //消息创建时间 （整型）
	MsgType      MsgType   `xml:"MsgType"`      //消息类型,文本:text 图片:image 语音: voice 视频:video 小视频:shortvideo 地理位置:location 链接:link
	Content      string    `xml:"Content"`      //文本消息内容
	PicURL       string    `xml:"PicUrl"`       //图片链接（由系统生成）
	MediaId      string    `xml:"MediaId"`      //消息媒体id，可以调用获取临时素材接口拉取数据。
	Format       string    `xml:"Format"`       //语音格式，如amr，speex等
	Recognition  string    `xml:"Recognition"`  //语音识别结果，UTF8编码
	ThumbMediaId string    `xml:"ThumbMediaId"` //视频消息缩略图的媒体id，可以调用多媒体文件下载接口拉取数据。
	LocationX    float64   `xml:"Location_X"`   //地理位置纬度
	LocationY    float64   `xml:"Location_Y"`   //地理位置经度
	Scale        int32     `xml:"Scale"`        //地图缩放大小
	Label        string    `xml:"Label"`        //地理位置信息
	Title        string    `xml:"Title"`        //消息标题
	Description  string    `xml:"Description"`  //消息描述
	URL          string    `xml:"Url"`          //消息链接
	MsgId        int64     `xml:"MsgId"`        //消息id
	Event        EventType `xml:"Event"`        //事件类型，subscribe(订阅)、unsubscribe(取消订阅)
	EventKey     string    `xml:"EventKey"`     //
	Ticket       string    `xml:"Ticket"`       //
	Latitude     float64   `xml:"latitude"`     // 地理位置纬度
	Longitude    float64   `xml:"longitude"`    // 地理位置经度
	Precision    float64   `xml:"precision"`    // 地理位置精度
}
