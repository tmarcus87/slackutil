package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/nlopes/slack"
	"os"
)

func main() {
	token, tokenOk := os.LookupEnv("SLACK_API_TOKEN")
	if !tokenOk {
		exit("SLACK_API_TOKEN is not given")
	}
	_, debug := os.LookupEnv("SLACK_DEBUG")

	if len(os.Args) <= 1 {
		exit("Must specify command")
	}

	client := slack.New(token, slack.OptionDebug(debug))

	switch command := os.Args[1]; command {
	case "post_message":
		postMessage(client, os.Args[2:])
	default:
		exit("Unknown command '" + command + "'")
	}

	os.Exit(0)
}

func postMessage(client *slack.Client, args []string) {
	f := flag.NewFlagSet("post_message", flag.ContinueOnError)
	ch := f.String("channel", "", "Channel")
	msg := f.String("message", "", "Post message")
	ts := f.String("thread", "", "Thread")

	if err := f.Parse(args); err != nil {
		exit("err:" + err.Error())
	}

	if *ch == "" {
		exit("'channel' parameter is not given")
	}
	if *msg == "" {
		exit("'message' parameter is not given")
	}

	opts :=
		[]slack.MsgOption{
			slack.MsgOptionUsername("CircleCI"),
			slack.MsgOptionText(*msg, false)}

	if *ts != "" {
		opts = append(opts, slack.MsgOptionTS(*ts))
	}

	client.SendMessage(*ch)

	respCh, respTs, respMsg, err := client.SendMessage(*ch, opts...)
	if err != nil {
		exit(err)
	}
	fmt.Println(j(map[string]string{
		"channel": respCh,
		"ts":      respTs,
		"message": respMsg,
	}))
}

func j(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		errbytes, _ := json.Marshal(err.Error())
		return string(errbytes)
	}
	return string(bytes)
}

func exit(v interface{}) {
	fmt.Println(j(v))
	os.Exit(1)
}
