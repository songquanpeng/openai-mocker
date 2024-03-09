package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func prompt2response(subject string) string {
	rand.Seed(time.Now().UnixNano())

	forwards := []string{
		"这个问题怎么又被顶上热搜了？？？\n以下为正文\n\n",
		"没想到会有这么多人看，还点赞？\n以下为正文\n\n",
		"没想到这么多人看，知乎小透明瑟瑟发抖。\n以下为正文\n\n",
		"喷子太多，评论区关了。\n以下为正文\n\n",
		"谢谢大家的支持，知乎小透明瑟瑟发抖。\n以下为正文\n\n",
		"一觉醒来一个赞都没有，知乎小透明瑟瑟发抖。\n以下为正文\n\n",
		"一觉醒来发现这么多赞，知乎小透明瑟瑟发抖。\n以下为正文\n\n",
		"",
	}

	openings := []string{
		"谢邀，",
		"泻药，",
		"谢不邀，怒答。",
		"蟹邀，",
		"怒答一发，",
		"不请自来，",
	}

	followOpenings := []string{
		"人在美国，刚下飞机。\n\n",
		"人在飞机，刚下美国。\n\n",
		"人在蒙古，刚下轮船。\n\n",
		"人在轮船，刚下蒙古。\n\n",
		"人在火星，刚下飞船。\n\n",
		"人在美帝，刚下飞机，正好看到这个问题。\n\n",
		"终于看到一个我能回答的问题。\n\n",
		"实名反对高赞答主的观点。\n\n",
		"圈子太小，熟人太多，匿了。\n\n",
		"怕被人认出来，匿了。\n\n",
		"过千赞就取匿。\n\n",
		"先说是不是，再问为什么\n\n",
		"@JustSong 已经回答得很好了，我来补充几句\n\n",
		"实名反对 @JustSong 的回答，大家不要被他带歪\n\n",
		"本人手机打字不易，请各位可以看完\n\n",
	}

	splitLines := []string{
		"---------分割线---------\n\n",
		"------------------------\n\n",
		"-------这是一条分割线-------\n\n",
		"-------华丽的分割线-------\n\n",
		"----分割线是这样画的吗----\n\n",
		"———手动分割———\n\n",
		"*************************\n\n",
		"~~~我是一条分割线~~~\n\n",
		"------FBI WARNING------\n\n",
		"•装•订•线•内•禁•止•作•答•否•则•记•零•分•\n\n",
		"→_→_→_→_→_→_→_→\n\n",
		"。。。。。。。。。。。。。。\n\n",
		"～～～～～～～～～～～～\n\n",
		"——好看的分割线(/ω＼)——\n\n",
		"～(￣▽￣～)~～(￣▽￣～)~～(￣▽￣～)~\n\n",
	}

	startStatement := []string{
		"要想正确看待这个问题，我们必须从历史的角度出发。\n\n",
		"你问我如何正确看待这个主题，那还能怎么看待，就像看待与主题类似的其他问题一样看待咯。\n\n",
		"客观地看，主题的存在是有其意义和价值的，我们必须用发展的眼光去看待主题。\n\n",
		"坦白的说，这个问题的存在还是有一定价值的，因为其牵涉到主题的内在性质。\n\n",
	}

	firstPoint := []string{
		"首先，主题之所以是主题，是必然有其自身的原因的，考虑到其内在特点，这也是自然而合理的。\n\n",
		"那首先呢，为什么会出现这样的问题呢？题主你有没有想过？为什么是主题而不是其他东西呢？其实第一个原因是是显而易见的，我就不点破了。\n\n",
		"第一点嘛，我们必须考虑到一个很关键的一点，即事物都是具有两面性的，那么主题也不例外，要想正确看待主题，我们必须明确看待主题的角度，考虑到这一点，其实也就很明了了嘛。\n\n",
	}

	secondPoint := []string{
		"其次，我们都知道，抛开剂量谈毒性，没什么意义，同样，如果我们不结合与主题相关的具体环境，空谈主题也是没什么价值的。\n\n",
		"另外，你自己好好想想，你所期望的主题是什么样子呢？希望大家都好好考虑一下主题的意义，我们就会发现，主题其实就应该这样。\n\n",
		"然后第二点，在主题相关的话题上，很多人总是尝试着抱着理性去思考主题，殊不知其内在的价值已经决定了主题之所以为主题，就在于主题本身的意义。\n\n",
		"然后呢，可能会有很多人误解主题的本意，看来大部分人在认识主题上还是太 naïve，作为圈内人，我还是希望大家多去了解一下于主题相关的资料，这样呢大家才能建立对主题更加客观实际的认知。\n\n",
	}

	endStatement := []string{
		"说了这么多，其实这个主题也没什么好说的，扯来扯去也就这些东西可讲。\n\n",
		"因此呢，这个主题它就是这个样子，我们就拿看待正常事物的眼光看待主题就好了。\n\n",
		"所以，就会出现像主题这样的情况。\n\n",
		"好了，废话不多说了，我认为出现像主题这种情况是不言而喻的。\n\n",
		"总之，这个主题大概就是这个样子，牵涉到部分人的利益，我不方便说太多，就这样。\n\n",
	}

	requestUp := []string{
		"大家不要只收藏，不点赞。",
		"收藏数是点赞数的两倍，这是知乎的风气么，i了i了。",
		"手机打字不易，喜欢的可以点个赞！",
		"手机打字不易，如果觉得对你有所帮助，麻烦点个赞。",
		"手机打字不易，给个赞再走呗？",
		"手机打字不易，各位看官点赞再走哦~",
		"手机打字不易，如果觉得对你有所帮助，麻烦点个赞。",
		"手机打字不易，求赞求感谢。",
	}

	requestReaders := []string{
		"最后打个广告，欢迎大家关注我的公众号「今天打代码了吗」，里面干货很多。\n\n",
		"打个广告，欢迎大家使用我的小程序「水文生成器」。\n\n",
		"欢迎大家关注我的公众号，等等，貌似我还没有公众号。\n\n",
		"最后自荐一下我的公众号「今天打代码了吗」，里面会不定期分享一些干货。\n\n",
		"最后自荐一下我的个人公众号「今天打代码了吗」，其里面的内容有一些与这个问题相关的干货哦。\n\n",
	}

	ends := []string{
		"以上。",
		"以上是我的看法。",
		"禁止转载。",
		"关闭评论区了，喷子太多。",
		"欢迎补充。",
		"以上是我的观点，如有遗漏错误，欢迎纠正。",
		"待补充。",
		"以上，发言完毕。",
		"随意转载。",
		"本回答遵循 CC-BY-SA-NC 协议，转载请注明出处。",
		"原创不易，转载请告知答主。",
		"第一次在知乎写回答，如有错漏请指出。",
		"最后，感谢你观看全文。不要忘了点个“赞+喜欢”！",
	}

	if strings.HasSuffix(subject, "?") || strings.HasSuffix(subject, "？") {
		subject = subject[:len(subject)-1]
	}
	var content string

	forward := randomChoose(forwards)
	if forward != "" {
		content += randomChoose(splitLines)
	}
	content += randomChoose(openings)
	content += randomChoose(followOpenings)

	statement := ""
	statement += randomChoose(startStatement)
	statement += randomChoose(firstPoint)
	statement += randomChoose(secondPoint)
	statement += randomChoose(endStatement)
	content += replaceSubject(statement, subject)
	content += randomChoose(splitLines)
	content += randomChoose(requestUp)
	content += randomChoose(requestReaders)
	content += randomChoose(ends)

	content += fmt.Sprintf("\n\n * 注意：以上回复仅供娱乐，由 https://github.com/songquanpeng/openai-mocker 生成*")
	return content
}

func randomChoose(options []string) string {
	if len(options) == 0 {
		return ""
	}
	return options[rand.Intn(len(options))]
}

func replaceSubject(statement string, subject string) string {
	return strings.ReplaceAll(statement, "主题", subject)
}
