package main

import (
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/slack-go/slack"
)

func main() {
	token := os.Getenv("SLACK_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")

	api := slack.New(token)

	for _, t := range buildMessages() {
		_, _, err := api.PostMessage(channel, t)
		if err != nil {
			panic(err)
		}
	}
}

func buildMessages() []slack.MsgOption {
	types := []slack.MsgOption{
		// Simple text
		slack.MsgOptionText("Hello world!", false),
		// With blocks
		// https://api.slack.com/reference/block-kit/blocks
		slack.MsgOptionBlocks(
			slack.NewContextBlock(
				uuid.New().String(),
				slack.NewTextBlockObject(
					slack.PlainTextType,
					"Context block",
					false,
					false,
				),
			),
			slack.NewSectionBlock(
				slack.NewTextBlockObject(
					slack.MarkdownType,
					"This is a mrkdwn section block :ghost: *this is bold*, and ~this is crossed out~, and <https://google.com|this is a link>",
					false,
					false,
				),
				nil,
				nil,
			),
			slack.NewImageBlock(
				"https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
				"Google",
				"google",
				&slack.TextBlockObject{
					Type: "plain_text",
					Text: "Google",
				},
			),
		),
		// With an attachment
		slack.MsgOptionAttachments(slack.Attachment{
			Title: "Attachment title",
			Text:  "Attachment text",
			Color: "#ff0000",
			Fields: []slack.AttachmentField{
				{
					Title: "Field title",
					Value: "Field value",
					Short: false,
				},
			},
			ThumbURL: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
		}),
		// With attachments
		slack.MsgOptionAttachments(
			slack.Attachment{
				Title: "1st attachment title",
				Text:  "1st attachment text",
				Color: "#00ff00",
				Fields: []slack.AttachmentField{
					{
						Title: "Field title",
						Value: "Field value",
						Short: true,
					},
				},
				ImageURL: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
			},
			slack.Attachment{
				Title: "2nd attachment title",
				Text:  "2nd attachment text",
				Color: "#00ff00",
				Fields: []slack.AttachmentField{
					{
						Title: "1st field title",
						Value: "1st field value",
						Short: true,
					},
					{
						Title: "2nd field title",
						Value: "2nd field value",
						Short: true,
					},
				},

				ThumbURL: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",

				// Invisible
				FromURL:     "https://www.google.com/",
				OriginalURL: "https://www.google.com/",

				// Appears at first
				AuthorName: "Google",
				AuthorLink: "https://www.google.com/",
				AuthorIcon: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",

				// Invisible
				ServiceName: "Google Service",
				ServiceIcon: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",

				Footer:     "Amazon",
				FooterIcon: "https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
			},
			slack.Attachment{
				// Title / Text は置いたら Blocks が使えなくなる
				Color: "#0000ff",
				Blocks: slack.Blocks{
					BlockSet: []slack.Block{
						slack.NewSectionBlock(
							slack.NewTextBlockObject(
								slack.MarkdownType,
								fmt.Sprintf("This is a markdown\nnow is ... `%s`", time.Now().String()),
								false,
								false,
							),
							nil,
							nil,
						),
						slack.NewContextBlock(
							uuid.New().String(),
							slack.NewTextBlockObject(
								slack.PlainTextType,
								"Context block",
								false,
								false,
							),
						),
						slack.NewHeaderBlock(
							slack.NewTextBlockObject(
								slack.PlainTextType,
								"Section block",
								false,
								false,
							),
						),
						slack.NewImageBlock(
							"https://www.google.com/images/branding/googlelogo/2x/googlelogo_color_272x92dp.png",
							"Google",
							"google",
							&slack.TextBlockObject{
								Type: "plain_text",
								Text: "Google",
							},
						),
					},
				},
			},
		),
	}
	return types
}
