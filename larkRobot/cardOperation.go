package larkrobot

import (
	"unsafe"

	"github.com/larksuite/botframework-go/SDK/message"
	"github.com/larksuite/botframework-go/SDK/protocol"
)

type CardOperation interface {
	buildTitle()
	buildBlock()
	buildConfig()
	buildButton()
}

type InteractiveCard struct {
	title         string //标题，"口语陪练测试工具(v0.1)""
	block         string //内容，"业务：EZ\n场景：情景对话"
	usid          []string
	builder       *message.CardBuilder
	cardOperation CardOperation
}

func (p *InteractiveCard) SetTitleContent(title string) {
	if p == nil {
		return
	}
	p.title = title
}

func (p *InteractiveCard) SetBlockContent(block string) {
	if p == nil {
		return
	}
	p.block = block
}

func (p *InteractiveCard) SetUsid(usid []string) {
	if p == nil {
		return
	}
	p.usid = usid
}

func (p *InteractiveCard) init() {
	if p == nil {
		return
	}
	p.builder = &message.CardBuilder{}
}

func (p *InteractiveCard) SetOperation(cardOperation *InteractiveCard) {
	if p == nil {
		return
	}
	p.cardOperation = cardOperation
}

func (p *InteractiveCard) Build() ([]byte, error) {
	if p == nil {
		return nil, nil
	}
	p.init()
	//改变建造方式设定
	p.SetOperation(p)

	p.cardOperation.buildTitle()
	p.cardOperation.buildConfig()
	p.cardOperation.buildBlock()
	p.cardOperation.buildButton()

	card, err := p.builder.Build()
	return card, err
}

func (p *InteractiveCard) buildTitle() {
	if p == nil {
		return
	}

	line := 1
	title := protocol.TextForm{
		Tag:     protocol.PLAIN_TEXT_E,
		Content: &p.title,
		Lines:   &line,
	}

	p.builder.AddHeader(title, "")
	p.builder.AddHRBlock() //设置分界线
}

func (p *InteractiveCard) buildBlock() {
	if p == nil {
		return
	}

	block := []protocol.FieldForm{*message.NewField(false, message.NewMDText(p.block, nil, nil, nil))}
	p.builder.AddDIVBlock(nil, block, nil)
}

func (p *InteractiveCard) buildConfig() {
	if p == nil {
		return
	}

	config := protocol.ConfigForm{
		MinVersion:     protocol.VersionForm{},
		WideScreenMode: true,
	}
	p.builder.SetConfig(config)
}

func (p *InteractiveCard) buildButton() {
	if p == nil || p.usid == nil {
		return
	}

	for _, text := range p.usid {
		payload := make(map[string]string, 0)
		payload["usid"] = text //KV对表示用户点击一个卡片后给后段POST的Header
		p.builder.AddActionBlock([]protocol.ActionElement{
			message.NewButton(message.NewMDText(text, nil, nil, nil),
				nil, nil, payload, protocol.PRIMARY, nil, "asyncButton"),
		})
	}
}

func Byte2Str(src []byte) string {
	return *(*string)(unsafe.Pointer(&src))
}

func Str2Byte(src string) []byte {
	return *(*[]byte)(unsafe.Pointer(&src))
}
