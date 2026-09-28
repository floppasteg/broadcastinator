# Config files

## Master config
When you fire up the program, it will need a "master configuration" file. This contains the URL that
the program will stream to, and where the next config file will be located.

The general structure looks like this: (this file and the complete structure, minus the can be found [here](../testdata/))

```conf
streamurl="rtmp://a.rtmp.youtube.com/live2/"
streamkey="streamkey.env"
service-spec="youtube"

program="/media/greypoupon/program/program.conf"
assets="/media/greypoupon/assets/assets.conf"

logfile="/mnt/d/temp/log.txt"
start="2029-06-24T12:30:00Z"
```

|Name|Description|
|----|-----------|
|`streamurl`|The URL pointing to an RTMP ingestion point|
|`streamkey`|Relative or absolute path to a file containing the stream key (note: depending on what `service-spec` is set to it may also be formatted or can contain just the key)|
|`service-spec`|Which service the program expects to stream to (most services have incompatible ways of sending video and sometimes caption data)|
|`program`|Absolute path to a program config|
|`assets`|Absolute path to an asset config|
|`logfile`|(optional) Absolute path to a file where (appends if the file already exists)|
|`start`|(optional) Time when the livestream will start (stream starts immediately once you hit the start button)|