slackutil
=========

```
# Send message
$ docker run --rm -e SLACK_API_TOKEN=token tmarcus87/slackutil \
  slackutil -channel '#name-of-channel' -message 'Hello!'
{"channel":"ChannelID", "ts": "ThreadID", "message": ""}

# Send message to thread
$ docker run --rm -e SLACK_API_TOKEN=token tmarcus87/slackutil \
  slackutil -channel '#name-of-channel' -message 'Hello!' -thread "THREAD_ID"
{"channel":"ChannelID", "ts": "ThreadID", "message": ""}
```