import { DownloadOutlined } from "@ant-design/icons/lib/icons"
import { Button, Card, Col, message, Row, Spin, Table, Tooltip } from "antd"
import "antd/dist/antd.min.css"
import Meta from "antd/lib/card/Meta"
import Search from "antd/lib/input/Search"
import Layout, { Content, Footer, Header } from "antd/lib/layout/layout"
import fileSize from "filesize"
import React from "react"
import { api } from "./Api"
import "./App.scss"

const columns = [
  {
    title: "iTag",
    dataIndex: "iTag",
    key: "iTag",
  },
  {
    title: "FPS",
    dataIndex: "fps",
    key: "fps",
  },
  {
    title: "Video Quality",
    dataIndex: "videoQuality",
  },
  {
    title: "Audio Quality",
    dataIndex: "audioQuality",
  },
  {
    title: "Audio Channels",
    dataIndex: "audioChannels",
  },
  {
    title: "Bitrate",
    dataIndex: "bitrate",
  },
  {
    title: "Format",
    dataIndex: "mimeType",
    render: (mimeType) => {
      // example: video/mp4; codecs="avc1.42001E, mp4a.40.2"
      let finalText = mimeType
      const types = mimeType.split(";", 2)
      if (types.length > 1) {
        finalText = types[0]
        const subTypes = types[0].split("/", 2)
        if (subTypes.length > 1) {
          finalText = subTypes[1].toUpperCase()
        }
      }
      return (
        <Tooltip title={mimeType}>
          <span>{finalText}</span>
        </Tooltip>
      )
    },
  },
  {
    title: "Size",
    dataIndex: "size",
    render: (bytes) => fileSize(bytes, { standard: "iec" }),
  },
  {
    title: "Download",
    dataIndex: "url",
    render: (text) => {
      if (text === "") {
        return null
      }
      return (
        <Button
          type="primary"
          shape="round"
          icon={<DownloadOutlined />}
          size="small"
          href={text}
          target="_blank"
        />
      )
    },
  },
]

class App extends React.Component {
  state = {
    url: "",
    data: null,
    loading: false,
  }

  loadYoutubeData = (url = "") => {
    if (url === "") {
      return
    }
    this.setState({ loading: true, url: url, data: null })
    api.youtube
      .get(url)
      .then((res) => {
        this.setState({ loading: false, data: res.data })
      })
      .catch((e) => {
        let errorMsg = e.message
        if (e.response && e.response.data) {
          errorMsg = e.response.data
        }
        message.error(errorMsg)
        this.setState({ loading: false })
      })
  }

  componentDidMount() {}

  rowIndicator = (record, _index) => {
    if (record.audioQuality !== "" && record.videoQuality !== "") {
      return "row-audio-video"
    } else if (record.audioQuality !== "") {
      return "row-audio"
    } else if (record.videoQuality !== "") {
      return "row-video"
    }
    return ""
  }

  render() {
    const { loading, data } = this.state

    let thumbnailContent = null
    let tableContent = null

    if (data !== null) {
      const formats = data.formats
      if (data.thumbnails && data.thumbnails.length > 0) {
        let tnl = data.thumbnails[0]
        let maxWidth = 0
        data.thumbnails.forEach((t) => {
          if (t.width < 480 && t.width > maxWidth) {
            maxWidth = t.width
            tnl = t
          }
        })
        thumbnailContent = (
          <Card style={{ width: tnl.width }} cover={<img alt={data.title} src={tnl.url} />}>
            <Meta title={`${data.title} [${data.duration}]`} description={data.author} />
          </Card>
        )
      }

      tableContent = (
        <>
          <Table
            dataSource={formats}
            columns={columns}
            pagination={false}
            size="small"
            rowClassName={this.rowIndicator}
          />
        </>
      )
    }

    return (
      <Layout>
        <Header>
          <span className="title">Youtube Video Download</span>
        </Header>
        <Content style={{ paddingTop: "10px" }}>
          <Row gutter={[0, 8]}>
            <Col span={12} offset={1}>
              <Spin spinning={loading} tip="In Progress..." size="large" style={{ marginTop: "60px" }}>
                <Search
                  placeholder="youtube url"
                  allowClear
                  enterButton="Get"
                  size="middle"
                  onSearch={this.loadYoutubeData}
                />
              </Spin>
            </Col>
            <Col span={22} offset={1}>
              {thumbnailContent}
            </Col>
            <Col span={22} offset={1}>
              {tableContent}
            </Col>
          </Row>
        </Content>
        <Footer />
      </Layout>
    )
  }
}

export default App
