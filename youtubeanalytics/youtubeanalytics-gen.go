// Package youtubeanalytics provides access to the YouTube Analytics API.
//
// See https://developers.google.com/youtube/analytics
//
// Usage example:
//
//   import "google.golang.org/api/youtubeanalytics/v2"
//   ...
//   youtubeanalyticsService, err := youtubeanalytics.New(oauthHttpClient)
package youtubeanalytics // import "google.golang.org/api/youtubeanalytics/v2"

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	context "golang.org/x/net/context"
	ctxhttp "golang.org/x/net/context/ctxhttp"
	gensupport "google.golang.org/api/gensupport"
	googleapi "google.golang.org/api/googleapi"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = ctxhttp.Do

const apiId = "youtubeAnalytics:v2"
const apiName = "youtubeAnalytics"
const apiVersion = "v2"
const basePath = "https://youtubeanalytics.googleapis.com/"

// OAuth2 scopes used by this API.
const (
	// Manage your YouTube account
	YoutubeScope = "https://www.googleapis.com/auth/youtube"

	// View your YouTube account
	YoutubeReadonlyScope = "https://www.googleapis.com/auth/youtube.readonly"

	// View and manage your assets and associated content on YouTube
	YoutubepartnerScope = "https://www.googleapis.com/auth/youtubepartner"

	// View monetary and non-monetary YouTube Analytics reports for your
	// YouTube content
	YtAnalyticsMonetaryReadonlyScope = "https://www.googleapis.com/auth/yt-analytics-monetary.readonly"

	// View YouTube Analytics reports for your YouTube content
	YtAnalyticsReadonlyScope = "https://www.googleapis.com/auth/yt-analytics.readonly"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.GroupItems = NewGroupItemsService(s)
	s.Groups = NewGroupsService(s)
	s.Reports = NewReportsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	GroupItems *GroupItemsService

	Groups *GroupsService

	Reports *ReportsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewGroupItemsService(s *Service) *GroupItemsService {
	rs := &GroupItemsService{s: s}
	return rs
}

type GroupItemsService struct {
	s *Service
}

func NewGroupsService(s *Service) *GroupsService {
	rs := &GroupsService{s: s}
	return rs
}

type GroupsService struct {
	s *Service
}

func NewReportsService(s *Service) *ReportsService {
	rs := &ReportsService{s: s}
	return rs
}

type ReportsService struct {
	s *Service
}

// EmptyResponse: Empty response.
type EmptyResponse struct {
	// Errors: Apiary error details
	Errors *Errors `json:"errors,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Errors") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *EmptyResponse) MarshalJSON() ([]byte, error) {
	type NoMethod EmptyResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ErrorProto: Describes one specific error.
type ErrorProto struct {
	// Argument: Error arguments, to be used when building user-friendly
	// error messages
	// given the error domain and code.  Different error codes require
	// different
	// arguments.
	Argument []string `json:"argument,omitempty"`

	// Code: Error code in the error domain. This should correspond to
	// a value of the enum type whose name is in domain. See
	// the core error domain in error_domain.proto.
	Code string `json:"code,omitempty"`

	// DebugInfo: Debugging information, which should not be
	// shared externally.
	DebugInfo string `json:"debugInfo,omitempty"`

	// Domain: Error domain. RoSy services can define their own
	// domain and error codes. This should normally be
	// the name of an enum type, such as: gdata.CoreErrorDomain
	Domain string `json:"domain,omitempty"`

	// ExternalErrorMessage: A short explanation for the error, which can be
	// shared outside Google.
	//
	// Please set domain, code and arguments whenever possible instead of
	// this
	// error message so that external APIs can build safe error
	// messages
	// themselves.
	//
	// External messages built in a RoSy interface will most likely refer
	// to
	// information and concepts that are not available externally and should
	// not
	// be exposed. It is safer if external APIs can understand the errors
	// and
	// decide what the error message should look like.
	ExternalErrorMessage string `json:"externalErrorMessage,omitempty"`

	// Location: Location of the error, as specified by the location
	// type.
	//
	// If location_type is PATH, this should be a path to a field
	// that's
	// relative to the request, using FieldPath
	// notation
	// (net/proto2/util/public/field_path.h).
	//
	// Examples:
	//   authenticated_user.gaia_id
	//   resource.address[2].country
	Location string `json:"location,omitempty"`

	// Possible values:
	//   "PATH" - location is an xpath-like path pointing
	// to the request field that caused the error.
	//   "OTHER" - other location type which can safely be
	// shared
	// externally.
	LocationType string `json:"locationType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Argument") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Argument") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ErrorProto) MarshalJSON() ([]byte, error) {
	type NoMethod ErrorProto
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Errors: Request Error information.
//
// The presence of an error field signals that the operation
// has failed.
type Errors struct {
	// Code: Global error code. Deprecated and ignored.
	// Set custom error codes in ErrorProto.domain and
	// ErrorProto.code
	// instead.
	//
	// Possible values:
	//   "BAD_REQUEST"
	//   "FORBIDDEN"
	//   "NOT_FOUND"
	//   "CONFLICT"
	//   "GONE"
	//   "PRECONDITION_FAILED"
	//   "INTERNAL_ERROR"
	//   "SERVICE_UNAVAILABLE"
	Code string `json:"code,omitempty"`

	// Error: Specific error description and codes
	Error []*ErrorProto `json:"error,omitempty"`

	// RequestId: Request identifier generated by the service, which can
	// be
	// used to identify the error in the logs
	RequestId string `json:"requestId,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Code") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Code") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *Errors) MarshalJSON() ([]byte, error) {
	type NoMethod Errors
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// Group: A group.
type Group struct {
	// ContentDetails: The `contentDetails` object contains additional
	// information about the
	// group, such as the number and type of items that it contains.
	ContentDetails *GroupContentDetails `json:"contentDetails,omitempty"`

	// Errors: Apiary error details
	Errors *Errors `json:"errors,omitempty"`

	// Etag: The Etag of this resource.
	Etag string `json:"etag,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the group.
	Id string `json:"id,omitempty"`

	// Kind: Identifies the API resource's type. The value will be
	// `youtube#group`.
	Kind string `json:"kind,omitempty"`

	// Snippet: The `snippet` object contains basic information about the
	// group, including
	// its creation date and name.
	Snippet *GroupSnippet `json:"snippet,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ContentDetails") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ContentDetails") to
	// include in API requests with the JSON null value. By default, fields
	// with empty values are omitted from API requests. However, any field
	// with an empty value appearing in NullFields will be sent to the
	// server as null. It is an error if a field in this list has a
	// non-empty value. This may be used to include null fields in Patch
	// requests.
	NullFields []string `json:"-"`
}

func (s *Group) MarshalJSON() ([]byte, error) {
	type NoMethod Group
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GroupContentDetails: A group's content details.
type GroupContentDetails struct {
	// ItemCount: The number of items in the group.
	ItemCount uint64 `json:"itemCount,omitempty,string"`

	// ItemType: The type of resources that the group contains.
	//
	// Valid values for this property are:
	//  * `youtube#channel`
	//  * `youtube#playlist`
	//  * `youtube#video`
	//  * `youtubePartner#asset`
	ItemType string `json:"itemType,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ItemCount") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ItemCount") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GroupContentDetails) MarshalJSON() ([]byte, error) {
	type NoMethod GroupContentDetails
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GroupItem: A group item.
type GroupItem struct {
	// Errors: Apiary error details
	Errors *Errors `json:"errors,omitempty"`

	// Etag: The Etag of this resource.
	Etag string `json:"etag,omitempty"`

	// GroupId: The ID that YouTube uses to uniquely identify the group that
	// contains the
	// item.
	GroupId string `json:"groupId,omitempty"`

	// Id: The ID that YouTube uses to uniquely identify the `channel`,
	// `video`,
	// `playlist`, or `asset` resource that is included in the group. Note
	// that
	// this ID refers specifically to the inclusion of that resource in
	// a
	// particular group and is different than the channel ID, video
	// ID,
	// playlist ID, or asset ID that uniquely identifies the resource
	// itself.
	// The `resource.id` property's value specifies the unique channel,
	// video,
	// playlist, or asset ID.
	Id string `json:"id,omitempty"`

	// Kind: Identifies the API resource's type. The value will be
	// `youtube#groupItem`.
	Kind string `json:"kind,omitempty"`

	// Resource: The `resource` object contains information that identifies
	// the item being
	// added to the group.
	Resource *GroupItemResource `json:"resource,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Errors") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GroupItem) MarshalJSON() ([]byte, error) {
	type NoMethod GroupItem
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

type GroupItemResource struct {
	// Id: The channel, video, playlist, or asset ID that YouTube uses to
	// uniquely
	// identify the item that is being added to the group.
	Id string `json:"id,omitempty"`

	// Kind: Identifies the type of resource being added to the
	// group.
	//
	// Valid values for this property are:
	//  * `youtube#channel`
	//  * `youtube#playlist`
	//  * `youtube#video`
	//  * `youtubePartner#asset`
	Kind string `json:"kind,omitempty"`

	// ForceSendFields is a list of field names (e.g. "Id") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Id") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GroupItemResource) MarshalJSON() ([]byte, error) {
	type NoMethod GroupItemResource
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// GroupSnippet: A group snippet.
type GroupSnippet struct {
	// PublishedAt: The date and time that the group was created. The value
	// is specified in
	// ISO 8601 (YYYY-MM-DDThh:mm:ss.sZ) format.
	PublishedAt string `json:"publishedAt,omitempty"`

	// Title: The group name. The value must be a non-empty string.
	Title string `json:"title,omitempty"`

	// ForceSendFields is a list of field names (e.g. "PublishedAt") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "PublishedAt") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *GroupSnippet) MarshalJSON() ([]byte, error) {
	type NoMethod GroupSnippet
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListGroupItemsResponse: Response message for
// GroupsService.ListGroupItems.
type ListGroupItemsResponse struct {
	// Errors: Apiary error details
	Errors *Errors `json:"errors,omitempty"`

	// Etag: The Etag of this resource.
	Etag string `json:"etag,omitempty"`

	// Items: A list of groups that match the API request parameters. Each
	// item in the
	// list represents a `groupItem` resource.
	Items []*GroupItem `json:"items,omitempty"`

	// Kind: Identifies the API resource's type. The value will
	// be
	// `youtube#groupItemListResponse`.
	Kind string `json:"kind,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Errors") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListGroupItemsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListGroupItemsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ListGroupsResponse: Response message for GroupsService.ListGroups.
type ListGroupsResponse struct {
	// Errors: Apiary error details
	Errors *Errors `json:"errors,omitempty"`

	// Etag: The Etag of this resource.
	Etag string `json:"etag,omitempty"`

	// Items: A list of groups that match the API request parameters. Each
	// item in the
	// list represents a `group` resource.
	Items []*Group `json:"items,omitempty"`

	// Kind: Identifies the API resource's type. The value will
	// be
	// `youtube#groupListResponse`.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The token that can be used as the value of the
	// `pageToken` parameter to
	// retrieve the next page in the result set.
	NextPageToken string `json:"nextPageToken,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Errors") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Errors") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ListGroupsResponse) MarshalJSON() ([]byte, error) {
	type NoMethod ListGroupsResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// QueryResponse: Response message for TargetedQueriesService.Query.
type QueryResponse struct {
	// ColumnHeaders: This value specifies information about the data
	// returned in the `rows`
	// fields. Each item in the `columnHeaders` list identifies a field
	// returned
	// in the `rows` value, which contains a list of comma-delimited data.
	// The
	// `columnHeaders` list will begin with the dimensions specified in the
	// API
	// request, which will be followed by the metrics specified in the
	// API
	// request. The order of both dimensions and metrics will match the
	// ordering
	// in the API request. For example, if the API request contains the
	// parameters
	// `dimensions=ageGroup,gender&metrics=viewerPercentage`, the API
	// response
	// will return columns in this order: `ageGroup`,
	// `gender`,
	// `viewerPercentage`.
	ColumnHeaders []*ResultTableColumnHeader `json:"columnHeaders,omitempty"`

	// Errors: When set, indicates that the operation failed.
	Errors *Errors `json:"errors,omitempty"`

	// Kind: This value specifies the type of data included in the API
	// response.
	// For the query method, the kind property value will
	// be
	// `youtubeAnalytics#resultTable`.
	Kind string `json:"kind,omitempty"`

	// Rows: The list contains all rows of the result table. Each item in
	// the list is
	// an array that contains comma-delimited data corresponding to a single
	// row
	// of data. The order of the comma-delimited data fields will match
	// the
	// order of the columns listed in the `columnHeaders` field.
	//
	// If no data is available for the given query, the `rows` element will
	// be
	// omitted from the response.
	//
	// The response for a query with the `day` dimension will not contain
	// rows for
	// the most recent days.
	Rows [][]interface{} `json:"rows,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "ColumnHeaders") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ColumnHeaders") to include
	// in API requests with the JSON null value. By default, fields with
	// empty values are omitted from API requests. However, any field with
	// an empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *QueryResponse) MarshalJSON() ([]byte, error) {
	type NoMethod QueryResponse
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// ResultTableColumnHeader: The description of a column of the result
// table.
type ResultTableColumnHeader struct {
	// ColumnType: The type of the column (`DIMENSION` or `METRIC`).
	ColumnType string `json:"columnType,omitempty"`

	// DataType: The type of the data in the column (`STRING`, `INTEGER`,
	// `FLOAT`, etc.).
	DataType string `json:"dataType,omitempty"`

	// Name: The name of the dimension or metric.
	Name string `json:"name,omitempty"`

	// ForceSendFields is a list of field names (e.g. "ColumnType") to
	// unconditionally include in API requests. By default, fields with
	// empty values are omitted from API requests. However, any non-pointer,
	// non-interface field appearing in ForceSendFields will be sent to the
	// server regardless of whether the field is empty or not. This may be
	// used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "ColumnType") to include in
	// API requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *ResultTableColumnHeader) MarshalJSON() ([]byte, error) {
	type NoMethod ResultTableColumnHeader
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "youtubeAnalytics.groupItems.delete":

type GroupItemsDeleteCall struct {
	s          *Service
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Removes an item from a group.
func (r *GroupItemsService) Delete() *GroupItemsDeleteCall {
	c := &GroupItemsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Id sets the optional parameter "id": The `id` parameter specifies the
// YouTube group item ID of the group item
// that is being deleted.
func (c *GroupItemsDeleteCall) Id(id string) *GroupItemsDeleteCall {
	c.urlParams_.Set("id", id)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": This parameter can only be used in a
// properly authorized request. **Note:**
// This parameter is intended exclusively for YouTube content partners
// that
// own and manage many different YouTube channels.
//
// The `onBehalfOfContentOwner` parameter indicates that the
// request's
// authorization credentials identify a YouTube user who is acting on
// behalf
// of the content owner specified in the parameter value. It allows
// content
// owners to authenticate once and get access to all their video and
// channel
// data, without having to provide authentication credentials for
// each
// individual channel. The account that the user authenticates with must
// be
// linked to the specified YouTube content owner.
func (c *GroupItemsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupItemsDeleteCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupItemsDeleteCall) Fields(s ...googleapi.Field) *GroupItemsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupItemsDeleteCall) Context(ctx context.Context) *GroupItemsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupItemsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupItemsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/groupItems")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubeAnalytics.groupItems.delete" call.
// Exactly one of *EmptyResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *EmptyResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *GroupItemsDeleteCall) Do(opts ...googleapi.CallOption) (*EmptyResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &EmptyResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Removes an item from a group.",
	//   "flatPath": "v2/groupItems",
	//   "httpMethod": "DELETE",
	//   "id": "youtubeAnalytics.groupItems.delete",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "id": {
	//       "description": "The `id` parameter specifies the YouTube group item ID of the group item\nthat is being deleted.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "This parameter can only be used in a properly authorized request. **Note:**\nThis parameter is intended exclusively for YouTube content partners that\nown and manage many different YouTube channels.\n\nThe `onBehalfOfContentOwner` parameter indicates that the request's\nauthorization credentials identify a YouTube user who is acting on behalf\nof the content owner specified in the parameter value. It allows content\nowners to authenticate once and get access to all their video and channel\ndata, without having to provide authentication credentials for each\nindividual channel. The account that the user authenticates with must be\nlinked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/groupItems",
	//   "response": {
	//     "$ref": "EmptyResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.groupItems.insert":

type GroupItemsInsertCall struct {
	s          *Service
	groupitem  *GroupItem
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Creates a group item.
func (r *GroupItemsService) Insert(groupitem *GroupItem) *GroupItemsInsertCall {
	c := &GroupItemsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.groupitem = groupitem
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": This parameter can only be used in a
// properly authorized request. **Note:**
// This parameter is intended exclusively for YouTube content partners
// that
// own and manage many different YouTube channels.
//
// The `onBehalfOfContentOwner` parameter indicates that the
// request's
// authorization credentials identify a YouTube user who is acting on
// behalf
// of the content owner specified in the parameter value. It allows
// content
// owners to authenticate once and get access to all their video and
// channel
// data, without having to provide authentication credentials for
// each
// individual channel. The account that the user authenticates with must
// be
// linked to the specified YouTube content owner.
func (c *GroupItemsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupItemsInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupItemsInsertCall) Fields(s ...googleapi.Field) *GroupItemsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupItemsInsertCall) Context(ctx context.Context) *GroupItemsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupItemsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupItemsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.groupitem)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/groupItems")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubeAnalytics.groupItems.insert" call.
// Exactly one of *GroupItem or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *GroupItem.ServerResponse.Header or (if a response was returned at
// all) in error.(*googleapi.Error).Header. Use googleapi.IsNotModified
// to check whether the returned error was because
// http.StatusNotModified was returned.
func (c *GroupItemsInsertCall) Do(opts ...googleapi.CallOption) (*GroupItem, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &GroupItem{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a group item.",
	//   "flatPath": "v2/groupItems",
	//   "httpMethod": "POST",
	//   "id": "youtubeAnalytics.groupItems.insert",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "This parameter can only be used in a properly authorized request. **Note:**\nThis parameter is intended exclusively for YouTube content partners that\nown and manage many different YouTube channels.\n\nThe `onBehalfOfContentOwner` parameter indicates that the request's\nauthorization credentials identify a YouTube user who is acting on behalf\nof the content owner specified in the parameter value. It allows content\nowners to authenticate once and get access to all their video and channel\ndata, without having to provide authentication credentials for each\nindividual channel. The account that the user authenticates with must be\nlinked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/groupItems",
	//   "request": {
	//     "$ref": "GroupItem"
	//   },
	//   "response": {
	//     "$ref": "GroupItem"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.groupItems.list":

type GroupItemsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns a collection of group items that match the API request
// parameters.
func (r *GroupItemsService) List() *GroupItemsListCall {
	c := &GroupItemsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// GroupId sets the optional parameter "groupId": The `groupId`
// parameter specifies the unique ID of the group for which you
// want to retrieve group items.
func (c *GroupItemsListCall) GroupId(groupId string) *GroupItemsListCall {
	c.urlParams_.Set("groupId", groupId)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": This parameter can only be used in a
// properly authorized request. **Note:**
// This parameter is intended exclusively for YouTube content partners
// that
// own and manage many different YouTube channels.
//
// The `onBehalfOfContentOwner` parameter indicates that the
// request's
// authorization credentials identify a YouTube user who is acting on
// behalf
// of the content owner specified in the parameter value. It allows
// content
// owners to authenticate once and get access to all their video and
// channel
// data, without having to provide authentication credentials for
// each
// individual channel. The account that the user authenticates with must
// be
// linked to the specified YouTube content owner.
func (c *GroupItemsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupItemsListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupItemsListCall) Fields(s ...googleapi.Field) *GroupItemsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *GroupItemsListCall) IfNoneMatch(entityTag string) *GroupItemsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupItemsListCall) Context(ctx context.Context) *GroupItemsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupItemsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupItemsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/groupItems")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubeAnalytics.groupItems.list" call.
// Exactly one of *ListGroupItemsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListGroupItemsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *GroupItemsListCall) Do(opts ...googleapi.CallOption) (*ListGroupItemsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListGroupItemsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a collection of group items that match the API request parameters.",
	//   "flatPath": "v2/groupItems",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.groupItems.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "groupId": {
	//       "description": "The `groupId` parameter specifies the unique ID of the group for which you\nwant to retrieve group items.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "This parameter can only be used in a properly authorized request. **Note:**\nThis parameter is intended exclusively for YouTube content partners that\nown and manage many different YouTube channels.\n\nThe `onBehalfOfContentOwner` parameter indicates that the request's\nauthorization credentials identify a YouTube user who is acting on behalf\nof the content owner specified in the parameter value. It allows content\nowners to authenticate once and get access to all their video and channel\ndata, without having to provide authentication credentials for each\nindividual channel. The account that the user authenticates with must be\nlinked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/groupItems",
	//   "response": {
	//     "$ref": "ListGroupItemsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.groups.delete":

type GroupsDeleteCall struct {
	s          *Service
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Delete: Deletes a group.
func (r *GroupsService) Delete() *GroupsDeleteCall {
	c := &GroupsDeleteCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Id sets the optional parameter "id": The `id` parameter specifies the
// YouTube group ID of the group that is
// being deleted.
func (c *GroupsDeleteCall) Id(id string) *GroupsDeleteCall {
	c.urlParams_.Set("id", id)
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": This parameter can only be used in a
// properly authorized request. **Note:**
// This parameter is intended exclusively for YouTube content partners
// that
// own and manage many different YouTube channels.
//
// The `onBehalfOfContentOwner` parameter indicates that the
// request's
// authorization credentials identify a YouTube user who is acting on
// behalf
// of the content owner specified in the parameter value. It allows
// content
// owners to authenticate once and get access to all their video and
// channel
// data, without having to provide authentication credentials for
// each
// individual channel. The account that the user authenticates with must
// be
// linked to the specified YouTube content owner.
func (c *GroupsDeleteCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsDeleteCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsDeleteCall) Fields(s ...googleapi.Field) *GroupsDeleteCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupsDeleteCall) Context(ctx context.Context) *GroupsDeleteCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupsDeleteCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupsDeleteCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/groups")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubeAnalytics.groups.delete" call.
// Exactly one of *EmptyResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *EmptyResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *GroupsDeleteCall) Do(opts ...googleapi.CallOption) (*EmptyResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &EmptyResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a group.",
	//   "flatPath": "v2/groups",
	//   "httpMethod": "DELETE",
	//   "id": "youtubeAnalytics.groups.delete",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "id": {
	//       "description": "The `id` parameter specifies the YouTube group ID of the group that is\nbeing deleted.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "This parameter can only be used in a properly authorized request. **Note:**\nThis parameter is intended exclusively for YouTube content partners that\nown and manage many different YouTube channels.\n\nThe `onBehalfOfContentOwner` parameter indicates that the request's\nauthorization credentials identify a YouTube user who is acting on behalf\nof the content owner specified in the parameter value. It allows content\nowners to authenticate once and get access to all their video and channel\ndata, without having to provide authentication credentials for each\nindividual channel. The account that the user authenticates with must be\nlinked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/groups",
	//   "response": {
	//     "$ref": "EmptyResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.groups.insert":

type GroupsInsertCall struct {
	s          *Service
	group      *Group
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Insert: Creates a group.
func (r *GroupsService) Insert(group *Group) *GroupsInsertCall {
	c := &GroupsInsertCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.group = group
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": This parameter can only be used in a
// properly authorized request. **Note:**
// This parameter is intended exclusively for YouTube content partners
// that
// own and manage many different YouTube channels.
//
// The `onBehalfOfContentOwner` parameter indicates that the
// request's
// authorization credentials identify a YouTube user who is acting on
// behalf
// of the content owner specified in the parameter value. It allows
// content
// owners to authenticate once and get access to all their video and
// channel
// data, without having to provide authentication credentials for
// each
// individual channel. The account that the user authenticates with must
// be
// linked to the specified YouTube content owner.
func (c *GroupsInsertCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsInsertCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsInsertCall) Fields(s ...googleapi.Field) *GroupsInsertCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupsInsertCall) Context(ctx context.Context) *GroupsInsertCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupsInsertCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupsInsertCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/groups")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubeAnalytics.groups.insert" call.
// Exactly one of *Group or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Group.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *GroupsInsertCall) Do(opts ...googleapi.CallOption) (*Group, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Group{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a group.",
	//   "flatPath": "v2/groups",
	//   "httpMethod": "POST",
	//   "id": "youtubeAnalytics.groups.insert",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "This parameter can only be used in a properly authorized request. **Note:**\nThis parameter is intended exclusively for YouTube content partners that\nown and manage many different YouTube channels.\n\nThe `onBehalfOfContentOwner` parameter indicates that the request's\nauthorization credentials identify a YouTube user who is acting on behalf\nof the content owner specified in the parameter value. It allows content\nowners to authenticate once and get access to all their video and channel\ndata, without having to provide authentication credentials for each\nindividual channel. The account that the user authenticates with must be\nlinked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/groups",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.groups.list":

type GroupsListCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// List: Returns a collection of groups that match the API request
// parameters. For
// example, you can retrieve all groups that the authenticated user
// owns,
// or you can retrieve one or more groups by their unique IDs.
func (r *GroupsService) List() *GroupsListCall {
	c := &GroupsListCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	return c
}

// Id sets the optional parameter "id": The `id` parameter specifies a
// comma-separated list of the YouTube group
// ID(s) for the resource(s) that are being retrieved. Each group must
// be
// owned by the authenticated user. In a `group` resource, the `id`
// property
// specifies the group's YouTube group ID.
//
// Note that if you do not specify a value for the `id` parameter, then
// you
// must set the `mine` parameter to `true`.
func (c *GroupsListCall) Id(id string) *GroupsListCall {
	c.urlParams_.Set("id", id)
	return c
}

// Mine sets the optional parameter "mine": This parameter can only be
// used in a properly authorized request. Set this
// parameter's value to true to retrieve all groups owned by the
// authenticated
// user.
func (c *GroupsListCall) Mine(mine bool) *GroupsListCall {
	c.urlParams_.Set("mine", fmt.Sprint(mine))
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": This parameter can only be used in a
// properly authorized request. **Note:**
// This parameter is intended exclusively for YouTube content partners
// that
// own and manage many different YouTube channels.
//
// The `onBehalfOfContentOwner` parameter indicates that the
// request's
// authorization credentials identify a YouTube user who is acting on
// behalf
// of the content owner specified in the parameter value. It allows
// content
// owners to authenticate once and get access to all their video and
// channel
// data, without having to provide authentication credentials for
// each
// individual channel. The account that the user authenticates with must
// be
// linked to the specified YouTube content owner.
func (c *GroupsListCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsListCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// PageToken sets the optional parameter "pageToken": The `pageToken`
// parameter identifies a specific page in the result set that
// should be returned. In an API response, the `nextPageToken`
// property
// identifies the next page that can be retrieved.
func (c *GroupsListCall) PageToken(pageToken string) *GroupsListCall {
	c.urlParams_.Set("pageToken", pageToken)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsListCall) Fields(s ...googleapi.Field) *GroupsListCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *GroupsListCall) IfNoneMatch(entityTag string) *GroupsListCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupsListCall) Context(ctx context.Context) *GroupsListCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupsListCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupsListCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/groups")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubeAnalytics.groups.list" call.
// Exactly one of *ListGroupsResponse or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *ListGroupsResponse.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *GroupsListCall) Do(opts ...googleapi.CallOption) (*ListGroupsResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &ListGroupsResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns a collection of groups that match the API request parameters. For\nexample, you can retrieve all groups that the authenticated user owns,\nor you can retrieve one or more groups by their unique IDs.",
	//   "flatPath": "v2/groups",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.groups.list",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "id": {
	//       "description": "The `id` parameter specifies a comma-separated list of the YouTube group\nID(s) for the resource(s) that are being retrieved. Each group must be\nowned by the authenticated user. In a `group` resource, the `id` property\nspecifies the group's YouTube group ID.\n\nNote that if you do not specify a value for the `id` parameter, then you\nmust set the `mine` parameter to `true`.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "mine": {
	//       "description": "This parameter can only be used in a properly authorized request. Set this\nparameter's value to true to retrieve all groups owned by the authenticated\nuser.",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "onBehalfOfContentOwner": {
	//       "description": "This parameter can only be used in a properly authorized request. **Note:**\nThis parameter is intended exclusively for YouTube content partners that\nown and manage many different YouTube channels.\n\nThe `onBehalfOfContentOwner` parameter indicates that the request's\nauthorization credentials identify a YouTube user who is acting on behalf\nof the content owner specified in the parameter value. It allows content\nowners to authenticate once and get access to all their video and channel\ndata, without having to provide authentication credentials for each\nindividual channel. The account that the user authenticates with must be\nlinked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "pageToken": {
	//       "description": "The `pageToken` parameter identifies a specific page in the result set that\nshould be returned. In an API response, the `nextPageToken` property\nidentifies the next page that can be retrieved.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/groups",
	//   "response": {
	//     "$ref": "ListGroupsResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// Pages invokes f for each page of results.
// A non-nil error returned from f will halt the iteration.
// The provided context supersedes any context provided to the Context method.
func (c *GroupsListCall) Pages(ctx context.Context, f func(*ListGroupsResponse) error) error {
	c.ctx_ = ctx
	defer c.PageToken(c.urlParams_.Get("pageToken")) // reset paging to original point
	for {
		x, err := c.Do()
		if err != nil {
			return err
		}
		if err := f(x); err != nil {
			return err
		}
		if x.NextPageToken == "" {
			return nil
		}
		c.PageToken(x.NextPageToken)
	}
}

// method id "youtubeAnalytics.groups.update":

type GroupsUpdateCall struct {
	s          *Service
	group      *Group
	urlParams_ gensupport.URLParams
	ctx_       context.Context
	header_    http.Header
}

// Update: Modifies a group. For example, you could change a group's
// title.
func (r *GroupsService) Update(group *Group) *GroupsUpdateCall {
	c := &GroupsUpdateCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.group = group
	return c
}

// OnBehalfOfContentOwner sets the optional parameter
// "onBehalfOfContentOwner": This parameter can only be used in a
// properly authorized request. **Note:**
// This parameter is intended exclusively for YouTube content partners
// that
// own and manage many different YouTube channels.
//
// The `onBehalfOfContentOwner` parameter indicates that the
// request's
// authorization credentials identify a YouTube user who is acting on
// behalf
// of the content owner specified in the parameter value. It allows
// content
// owners to authenticate once and get access to all their video and
// channel
// data, without having to provide authentication credentials for
// each
// individual channel. The account that the user authenticates with must
// be
// linked to the specified YouTube content owner.
func (c *GroupsUpdateCall) OnBehalfOfContentOwner(onBehalfOfContentOwner string) *GroupsUpdateCall {
	c.urlParams_.Set("onBehalfOfContentOwner", onBehalfOfContentOwner)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *GroupsUpdateCall) Fields(s ...googleapi.Field) *GroupsUpdateCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *GroupsUpdateCall) Context(ctx context.Context) *GroupsUpdateCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *GroupsUpdateCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *GroupsUpdateCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.group)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/groups")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubeAnalytics.groups.update" call.
// Exactly one of *Group or error will be non-nil. Any non-2xx status
// code is an error. Response headers are in either
// *Group.ServerResponse.Header or (if a response was returned at all)
// in error.(*googleapi.Error).Header. Use googleapi.IsNotModified to
// check whether the returned error was because http.StatusNotModified
// was returned.
func (c *GroupsUpdateCall) Do(opts ...googleapi.CallOption) (*Group, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &Group{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Modifies a group. For example, you could change a group's title.",
	//   "flatPath": "v2/groups",
	//   "httpMethod": "PUT",
	//   "id": "youtubeAnalytics.groups.update",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "onBehalfOfContentOwner": {
	//       "description": "This parameter can only be used in a properly authorized request. **Note:**\nThis parameter is intended exclusively for YouTube content partners that\nown and manage many different YouTube channels.\n\nThe `onBehalfOfContentOwner` parameter indicates that the request's\nauthorization credentials identify a YouTube user who is acting on behalf\nof the content owner specified in the parameter value. It allows content\nowners to authenticate once and get access to all their video and channel\ndata, without having to provide authentication credentials for each\nindividual channel. The account that the user authenticates with must be\nlinked to the specified YouTube content owner.",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v2/groups",
	//   "request": {
	//     "$ref": "Group"
	//   },
	//   "response": {
	//     "$ref": "Group"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}

// method id "youtubeAnalytics.reports.query":

type ReportsQueryCall struct {
	s            *Service
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// Query: Retrieve your YouTube Analytics reports.
func (r *ReportsService) Query(ids string, startDate string, endDate string, metrics string, dimensions string) *ReportsQueryCall {
	c := &ReportsQueryCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.urlParams_.Set("ids", ids)
	c.urlParams_.Set("start-date", startDate)
	c.urlParams_.Set("end-date", endDate)
	c.urlParams_.Set("metrics", metrics)
	c.urlParams_.Set("dimensions", dimensions)
	return c
}

// Currency sets the optional parameter "currency": The currency to
// which financial metrics should be converted. The default is
// US Dollar (USD). If the result contains no financial metrics, this
// flag
// will be ignored. Responds with an error if the specified currency is
// not
// recognized.",
// pattern: [A-Z]{3}
func (c *ReportsQueryCall) Currency(currency string) *ReportsQueryCall {
	c.urlParams_.Set("currency", currency)
	return c
}

// Dimensions sets the optional parameter "dimensions": A
// comma-separated list of YouTube Analytics dimensions, such as `views`
// or
// `ageGroup,gender`. See the
// [Available
// Reports](/youtube/analytics/v2/available_reports) document for a list
// of
// the reports that you can retrieve and the dimensions used for
// those
// reports. Also see the
// [Dimensions](/youtube/analytics/v2/dimsmets/dims)
// document for definitions of those dimensions."
// pattern: [0-9a-zA-Z,]+
func (c *ReportsQueryCall) Dimensions(dimensions string) *ReportsQueryCall {
	c.urlParams_.Set("dimensions", dimensions)
	return c
}

// EndDate sets the optional parameter "endDate": The end date for
// fetching YouTube Analytics data. The value should be in
// `YYYY-MM-DD` format.
// required: true, pattern: [0-9]{4}-[0-9]{2}-[0-9]{2}
func (c *ReportsQueryCall) EndDate(endDate string) *ReportsQueryCall {
	c.urlParams_.Set("endDate", endDate)
	return c
}

// Filters sets the optional parameter "filters": A list of filters that
// should be applied when retrieving YouTube Analytics
// data. The [Available
// Reports](/youtube/analytics/v2/available_reports)
// document identifies the dimensions that can be used to filter each
// report,
// and the [Dimensions](/youtube/analytics/v2/dimsmets/dims)  document
// defines
// those dimensions. If a request uses multiple filters, join them
// together
// with a semicolon (`;`), and the returned result table will satisfy
// both
// filters. For example, a filters parameter value
// of
// `video==dMH0bHeiRNg;country==IT` restricts the result set to include
// data
// for the given video in Italy.",
func (c *ReportsQueryCall) Filters(filters string) *ReportsQueryCall {
	c.urlParams_.Set("filters", filters)
	return c
}

// Ids sets the optional parameter "ids": Identifies the YouTube channel
// or content owner for which you are
// retrieving YouTube Analytics data.
//
// - To request data for a YouTube user, set the `ids` parameter value
// to
//   `channel==CHANNEL_ID`, where `CHANNEL_ID` specifies the unique
// YouTube
//   channel ID.
// - To request data for a YouTube CMS content owner, set the `ids`
// parameter
//   value to `contentOwner==OWNER_NAME`, where `OWNER_NAME` is the CMS
// name
//   of the content owner.
// required: true, pattern: [a-zA-Z]+==[a-zA-Z0-9_+-]+
func (c *ReportsQueryCall) Ids(ids string) *ReportsQueryCall {
	c.urlParams_.Set("ids", ids)
	return c
}

// IncludeHistoricalChannelData sets the optional parameter
// "includeHistoricalChannelData": If set to true historical data (i.e.
// channel data from before the linking
// of the channel to the content owner) will be retrieved.",
func (c *ReportsQueryCall) IncludeHistoricalChannelData(includeHistoricalChannelData bool) *ReportsQueryCall {
	c.urlParams_.Set("includeHistoricalChannelData", fmt.Sprint(includeHistoricalChannelData))
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of rows to include in the response.",
// minValue: 1
func (c *ReportsQueryCall) MaxResults(maxResults int64) *ReportsQueryCall {
	c.urlParams_.Set("maxResults", fmt.Sprint(maxResults))
	return c
}

// Metrics sets the optional parameter "metrics": A comma-separated list
// of YouTube Analytics metrics, such as `views` or
// `likes,dislikes`. See the
// [Available Reports](/youtube/analytics/v2/available_reports)
// document for
// a list of the reports that you can retrieve and the metrics
// available in each report, and see
// the
// [Metrics](/youtube/analytics/v2/dimsmets/mets) document for
// definitions of
// those metrics.
// required: true, pattern: [0-9a-zA-Z,]+
func (c *ReportsQueryCall) Metrics(metrics string) *ReportsQueryCall {
	c.urlParams_.Set("metrics", metrics)
	return c
}

// Sort sets the optional parameter "sort": A comma-separated list of
// dimensions or metrics that determine the sort
// order for YouTube Analytics data. By default the sort order is
// ascending.
// The '`-`' prefix causes descending sort order.",
// pattern: [-0-9a-zA-Z,]+
func (c *ReportsQueryCall) Sort(sort string) *ReportsQueryCall {
	c.urlParams_.Set("sort", sort)
	return c
}

// StartDate sets the optional parameter "startDate": The start date for
// fetching YouTube Analytics data. The value should be in
// `YYYY-MM-DD` format.
// required: true, pattern: "[0-9]{4}-[0-9]{2}-[0-9]{2}
func (c *ReportsQueryCall) StartDate(startDate string) *ReportsQueryCall {
	c.urlParams_.Set("startDate", startDate)
	return c
}

// StartIndex sets the optional parameter "startIndex": An index of the
// first entity to retrieve. Use this parameter as a
// pagination mechanism along with the max-results parameter
// (one-based,
// inclusive).",
// minValue: 1
func (c *ReportsQueryCall) StartIndex(startIndex int64) *ReportsQueryCall {
	c.urlParams_.Set("startIndex", fmt.Sprint(startIndex))
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *ReportsQueryCall) Fields(s ...googleapi.Field) *ReportsQueryCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *ReportsQueryCall) IfNoneMatch(entityTag string) *ReportsQueryCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *ReportsQueryCall) Context(ctx context.Context) *ReportsQueryCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *ReportsQueryCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *ReportsQueryCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	urls := googleapi.ResolveRelative(c.s.BasePath, "v2/reports")
	urls += "?" + c.urlParams_.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	req.Header = reqHeaders
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "youtubeAnalytics.reports.query" call.
// Exactly one of *QueryResponse or error will be non-nil. Any non-2xx
// status code is an error. Response headers are in either
// *QueryResponse.ServerResponse.Header or (if a response was returned
// at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *ReportsQueryCall) Do(opts ...googleapi.CallOption) (*QueryResponse, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &QueryResponse{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieve your YouTube Analytics reports.",
	//   "flatPath": "v2/reports",
	//   "httpMethod": "GET",
	//   "id": "youtubeAnalytics.reports.query",
	//   "parameterOrder": [],
	//   "parameters": {
	//     "currency": {
	//       "description": "The currency to which financial metrics should be converted. The default is\nUS Dollar (USD). If the result contains no financial metrics, this flag\nwill be ignored. Responds with an error if the specified currency is not\nrecognized.\",\npattern: [A-Z]{3}",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "dimensions": {
	//       "description": "A comma-separated list of YouTube Analytics dimensions, such as `views` or\n`ageGroup,gender`. See the [Available\nReports](/youtube/analytics/v2/available_reports) document for a list of\nthe reports that you can retrieve and the dimensions used for those\nreports. Also see the [Dimensions](/youtube/analytics/v2/dimsmets/dims)\ndocument for definitions of those dimensions.\"\npattern: [0-9a-zA-Z,]+",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "endDate": {
	//       "description": "The end date for fetching YouTube Analytics data. The value should be in\n`YYYY-MM-DD` format.\nrequired: true, pattern: [0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "filters": {
	//       "description": "A list of filters that should be applied when retrieving YouTube Analytics\ndata. The [Available Reports](/youtube/analytics/v2/available_reports)\ndocument identifies the dimensions that can be used to filter each report,\nand the [Dimensions](/youtube/analytics/v2/dimsmets/dims)  document defines\nthose dimensions. If a request uses multiple filters, join them together\nwith a semicolon (`;`), and the returned result table will satisfy both\nfilters. For example, a filters parameter value of\n`video==dMH0bHeiRNg;country==IT` restricts the result set to include data\nfor the given video in Italy.\",",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "ids": {
	//       "description": "Identifies the YouTube channel or content owner for which you are\nretrieving YouTube Analytics data.\n\n- To request data for a YouTube user, set the `ids` parameter value to\n  `channel==CHANNEL_ID`, where `CHANNEL_ID` specifies the unique YouTube\n  channel ID.\n- To request data for a YouTube CMS content owner, set the `ids` parameter\n  value to `contentOwner==OWNER_NAME`, where `OWNER_NAME` is the CMS name\n  of the content owner.\nrequired: true, pattern: [a-zA-Z]+==[a-zA-Z0-9_+-]+",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "includeHistoricalChannelData": {
	//       "description": "If set to true historical data (i.e. channel data from before the linking\nof the channel to the content owner) will be retrieved.\",",
	//       "location": "query",
	//       "type": "boolean"
	//     },
	//     "maxResults": {
	//       "description": "The maximum number of rows to include in the response.\",\nminValue: 1",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "metrics": {
	//       "description": "A comma-separated list of YouTube Analytics metrics, such as `views` or\n`likes,dislikes`. See the\n[Available Reports](/youtube/analytics/v2/available_reports)  document for\na list of the reports that you can retrieve and the metrics\navailable in each report, and see the\n[Metrics](/youtube/analytics/v2/dimsmets/mets) document for definitions of\nthose metrics.\nrequired: true, pattern: [0-9a-zA-Z,]+",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "sort": {
	//       "description": "A comma-separated list of dimensions or metrics that determine the sort\norder for YouTube Analytics data. By default the sort order is ascending.\nThe '`-`' prefix causes descending sort order.\",\npattern: [-0-9a-zA-Z,]+",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startDate": {
	//       "description": "The start date for fetching YouTube Analytics data. The value should be in\n`YYYY-MM-DD` format.\nrequired: true, pattern: \"[0-9]{4}-[0-9]{2}-[0-9]{2}",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "startIndex": {
	//       "description": "An index of the first entity to retrieve. Use this parameter as a\npagination mechanism along with the max-results parameter (one-based,\ninclusive).\",\nminValue: 1",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     }
	//   },
	//   "path": "v2/reports",
	//   "response": {
	//     "$ref": "QueryResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/youtube",
	//     "https://www.googleapis.com/auth/youtube.readonly",
	//     "https://www.googleapis.com/auth/youtubepartner",
	//     "https://www.googleapis.com/auth/yt-analytics-monetary.readonly",
	//     "https://www.googleapis.com/auth/yt-analytics.readonly"
	//   ]
	// }

}
