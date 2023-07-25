<?xml version="1.0" encoding="UTF-8"?>
<rss version="2.0" xmlns:itunes="http://www.itunes.com/dtds/podcast-1.0.dtd">
<channel>
  <title>{{.Name }}</title>
    <itunes:author>{{ .Author }}</itunes:author>
  <description>{{ .Description }}</description>
    <itunes:image href="{{ .Cover }}"/>
  <language>{{ .Language }}</language>
  <link>{{ .Link }}</link>
{{with .Episodes }}
{{ range . }}
<item>
      <title>{{ .Title }}</title>
      <description>{{ .Description }}</description>
      <pubDate>{{ .PubDate.Format  "Mon, 02 Jan 2006 15:04:06 -0700" }}</pubDate>
      <enclosure url="{{ .Link }}"
                 type="audio/mpeg" length="{{ .Length }}"/>
</item>
{{ end }}
{{ end }}
</channel>
</rss>
