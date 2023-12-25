package openhab

import "github.com/jinzhu/copier"

// rest/things?staticDataOnly=false
type Things []struct {
	Channels      []Channels    `json:"channels"`
	StatusInfo    StatusInfo    `json:"statusInfo"`
	Editable      bool          `json:"editable"`
	Label         string        `json:"label"`
	Configuration Configuration `json:"configuration,omitempty"`
	Properties    Properties    `json:"properties,omitempty"`
	UID           string        `json:"UID"`
	ThingTypeUID  string        `json:"thingTypeUID"`
	BridgeUID     string        `json:"bridgeUID,omitempty"`
}

// Clone return copy
func (t *Things) Clone() *Things {
	c := &Things{}
	copier.Copy(&c, &t)
	return c
}

type Channels struct {
	LinkedItems    []any         `json:"linkedItems"`
	UID            string        `json:"uid"`
	ID             string        `json:"id"`
	ChannelTypeUID string        `json:"channelTypeUID"`
	Kind           string        `json:"kind"`
	Label          string        `json:"label"`
	Description    string        `json:"description"`
	DefaultTags    []any         `json:"defaultTags"`
	Properties     Properties    `json:"properties"`
	Configuration  Configuration `json:"configuration"`
	ItemType       string        `json:"itemType,omitempty"`
}

// Clone return copy
func (t *Channels) Clone() *Channels {
	c := &Channels{}
	copier.Copy(&c, &t)
	return c
}

type StatusInfo struct {
	Status       string `json:"status"`
	StatusDetail string `json:"statusDetail"`
}

// Clone return copy
func (t *StatusInfo) Clone() *StatusInfo {
	c := &StatusInfo{}
	copier.Copy(&c, &t)
	return c
}

type Properties struct {
	DeviceIP             string `json:"deviceIp"`
	BrightnessAutoOn     bool   `json:"brightnessAutoOn"`
	UpdateInterval       int    `json:"updateInterval"`
	DeviceMode           string `json:"deviceMode"`
	DeviceType           string `json:"deviceType"`
	MacAddress           string `json:"macAddress"`
	FirmwareVersion      string `json:"firmwareVersion"`
	WifiNetwork          string `json:"wifiNetwork"`
	DeviceGeneration     string `json:"deviceGeneration"`
	CoiotAutoEnable      string `json:"coiotAutoEnable"`
	UpdateNewVersion     string `json:"updateNewVersion"`
	Vendor               string `json:"vendor"`
	ServiceName          string `json:"serviceName"`
	UpdateCurrentVersion string `json:"updateCurrentVersion"`
	DeviceName           string `json:"deviceName"`
	DevUpdatePeriod      string `json:"devUpdatePeriod"`
	UpdateStatus         string `json:"updateStatus"`
	ModelID              string `json:"modelId"`
	DeviceAuth           string `json:"deviceAuth"`
	UpdateAvailable      string `json:"updateAvailable"`
	CoapDeviceDescr      string `json:"coapDeviceDescr"`
	DeviceHwRev          string `json:"deviceHwRev"`
	CoapVersion          string `json:"coapVersion"`
	DeviceHwBatch        string `json:"deviceHwBatch"`
	PtzHome              string `json:"ptz_home"`
	PtzFocus             string `json:"ptz_focus"`
	Resolution           string `json:"resolution"`
	Ptz                  string `json:"ptz"`
	PtzAbs               string `json:"ptz_abs"`
	Host                 string `json:"host"`
	PtzZoom              string `json:"ptz_zoom"`
	PtzSpeed             string `json:"ptz_speed"`
	DeviceID             string `json:"deviceID"`
	PtzIris              string `json:"ptz_iris"`
	Model                string `json:"model"`
	Type                 string `json:"type"`
	PtzPan               string `json:"ptz_pan"`
	PtzZoomSpeed         string `json:"ptz_zoom_speed"`
	PtzTilt              string `json:"ptz_tilt"`
	PtzAutofocus         string `json:"ptz_autofocus"`
	NumberMeters         string `json:"numberMeters"`
	NumberRelays         string `json:"numberRelays"`
	NumberRollers        string `json:"numberRollers"`
}

// Clone return copy
func (t *Properties) Clone() *Properties {
	c := &Properties{}
	copier.Copy(&c, &t)
	return c
}

type Configuration struct {
	EventsCoIoT         bool   `json:"eventsCoIoT"`
	DeviceIP            string `json:"deviceIp"`
	UpdateInterval      int    `json:"updateInterval"`
	EventsSensorReport  bool   `json:"eventsSensorReport"`
	LowBattery          int    `json:"lowBattery"`
	RefreshRateEvents   int    `json:"refreshRateEvents"`
	SnapshotStreamID    int    `json:"snapshotStreamId"`
	RefreshRateMdParam  int    `json:"refreshRateMdParam"`
	RefreshRateSnapshot int    `json:"refreshRateSnapshot"`
	EnableBluGateway    bool   `json:"enableBluGateway"`
	LwtQos              int    `json:"lwtQos"`
	Publickeypin        bool   `json:"publickeypin"`
	Clientid            string `json:"clientid"`
	KeepAlive           int    `json:"keepAlive"`
	HostnameValidated   bool   `json:"hostnameValidated"`
	BirthRetain         bool   `json:"birthRetain"`
	Secure              bool   `json:"secure"`
	Certificatepin      bool   `json:"certificatepin"`
	ShutdownRetain      bool   `json:"shutdownRetain"`
	Protocol            string `json:"protocol"`
	Qos                 int    `json:"qos"`
	ReconnectTime       int    `json:"reconnectTime"`
	MqttVersion         string `json:"mqttVersion"`
	Host                string `json:"host"`
	LwtRetain           bool   `json:"lwtRetain"`
	EnableDiscovery     bool   `json:"enableDiscovery"`
	Apikey              string `json:"apikey"`
	Sound               string `json:"sound"`
	Expire              int    `json:"expire"`
	Format              string `json:"format"`
	Title               string `json:"title"`
	User                string `json:"user"`
	Retry               int    `json:"retry"`
	Timeout             int    `json:"timeout"`
	AcceptSsl           bool   `json:"acceptSsl"`
	Password            string `json:"password"`
	Port                int    `json:"port"`
	Username            string `json:"username"`
}

// Clone return copy
func (t *Configuration) Clone() *Configuration {
	c := &Configuration{}
	copier.Copy(&c, &t)
	return c
}

type ErrorNotFound struct {
}

func (t *ErrorNotFound) Error() string {
	return "Not Found"
}
