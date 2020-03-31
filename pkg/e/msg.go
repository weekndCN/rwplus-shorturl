package e

import (
    "strings"
)

// 设置错误返回信息Map
var MsgFlags = map[int]string {
    SUCCESS                         :"成功",
    ERROR                           :"失败",
    INVALID_PARAMS                  :"请求参数错误",
}

var ValidFlags = map[string]string {
    "Required"      :"不能为空",
    "Min"           :"最小值 为 %d",
    "Max"           :"最大值 为 %d",
    "Range"         :"范围 为 %d 到 %d",
    "MinSize"       :"最短长度 为 %d",
    "MaxSize"       :"最大长度 为 %d",
    "Length"        :"长度必须 为 %d",
    "Alpha"         :"必须是有效的字母",
    "Numeric"       :"必须是有效的数字",
    "AlphaNumeric"  :"必须是有效的字母或数字",
    "Match"         :"必须匹配 %s",
    "NoMatch"       :"必须不匹配 %s",
    "AlphaDash"     :"必须是有效的字母、数字或连接符号(-_)",
    "Email"         :"必须是有效的电子邮件地址",
    "IP"            :"必须是有效的IP地址",
    "Base64"        :"必须是有效的base64字符",
    "Mobile"        :"必须是有效的手机号码",
    "Tel"           :"必须是有效的电话号码",
    "Phone"         :"必须是有效的电话或移动电话号码",
    "ZipCode"       :"必须是有效的邮政编码",
}

// 返回Flags信息
func GetMsg(code int) string {
    msg, ok := MsgFlags[code]
    if ok {
        return msg
    }

    return MsgFlags[ERROR]
}


func ValidMsg(errkey string) string {
    arr := strings.Split(errkey, ".")
    msg, ok := ValidFlags[arr[1]]
    if ok {
        return arr[0]+msg
    }

    return arr[0]+ValidFlags["Required"]
}
