# Go API client for overseerr_go

This is the documentation for the Overseerr API backend.

Two primary authentication methods are supported:

- **Cookie Authentication**: A valid sign-in to the `/auth/plex` or `/auth/local` will generate a valid authentication cookie.
- **API Key Authentication**: Sign-in is also possible by passing an `X-Api-Key` header along with a valid API Key generated by Overseerr.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import overseerr_go "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `overseerr_go.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), overseerr_go.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `overseerr_go.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), overseerr_go.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `overseerr_go.ContextOperationServerIndices` and `overseerr_go.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), overseerr_go.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), overseerr_go.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:5055/api/v1*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AuthAPI* | [**AuthLocalPost**](docs/AuthAPI.md#authlocalpost) | **Post** /auth/local | Sign in using a local account
*AuthAPI* | [**AuthLogoutPost**](docs/AuthAPI.md#authlogoutpost) | **Post** /auth/logout | Sign out and clear session cookie
*AuthAPI* | [**AuthMeGet**](docs/AuthAPI.md#authmeget) | **Get** /auth/me | Get logged-in user
*AuthAPI* | [**AuthPlexPost**](docs/AuthAPI.md#authplexpost) | **Post** /auth/plex | Sign in using a Plex token
*CollectionAPI* | [**CollectionCollectionIdGet**](docs/CollectionAPI.md#collectioncollectionidget) | **Get** /collection/{collectionId} | Get collection details
*IssueAPI* | [**IssueCommentCommentIdDelete**](docs/IssueAPI.md#issuecommentcommentiddelete) | **Delete** /issueComment/{commentId} | Delete issue comment
*IssueAPI* | [**IssueCommentCommentIdGet**](docs/IssueAPI.md#issuecommentcommentidget) | **Get** /issueComment/{commentId} | Get issue comment
*IssueAPI* | [**IssueCommentCommentIdPut**](docs/IssueAPI.md#issuecommentcommentidput) | **Put** /issueComment/{commentId} | Update issue comment
*IssueAPI* | [**IssueCountGet**](docs/IssueAPI.md#issuecountget) | **Get** /issue/count | Gets issue counts
*IssueAPI* | [**IssueGet**](docs/IssueAPI.md#issueget) | **Get** /issue | Get all issues
*IssueAPI* | [**IssueIssueIdCommentPost**](docs/IssueAPI.md#issueissueidcommentpost) | **Post** /issue/{issueId}/comment | Create a comment
*IssueAPI* | [**IssueIssueIdDelete**](docs/IssueAPI.md#issueissueiddelete) | **Delete** /issue/{issueId} | Delete issue
*IssueAPI* | [**IssueIssueIdGet**](docs/IssueAPI.md#issueissueidget) | **Get** /issue/{issueId} | Get issue
*IssueAPI* | [**IssueIssueIdStatusPost**](docs/IssueAPI.md#issueissueidstatuspost) | **Post** /issue/{issueId}/{status} | Update an issue&#39;s status
*IssueAPI* | [**IssuePost**](docs/IssueAPI.md#issuepost) | **Post** /issue | Create new issue
*MediaAPI* | [**MediaGet**](docs/MediaAPI.md#mediaget) | **Get** /media | Get media
*MediaAPI* | [**MediaMediaIdDelete**](docs/MediaAPI.md#mediamediaiddelete) | **Delete** /media/{mediaId} | Delete media item
*MediaAPI* | [**MediaMediaIdStatusPost**](docs/MediaAPI.md#mediamediaidstatuspost) | **Post** /media/{mediaId}/{status} | Update media status
*MediaAPI* | [**MediaMediaIdWatchDataGet**](docs/MediaAPI.md#mediamediaidwatchdataget) | **Get** /media/{mediaId}/watch_data | Get watch data
*MoviesAPI* | [**MovieMovieIdGet**](docs/MoviesAPI.md#moviemovieidget) | **Get** /movie/{movieId} | Get movie details
*MoviesAPI* | [**MovieMovieIdRatingsGet**](docs/MoviesAPI.md#moviemovieidratingsget) | **Get** /movie/{movieId}/ratings | Get movie ratings
*MoviesAPI* | [**MovieMovieIdRatingscombinedGet**](docs/MoviesAPI.md#moviemovieidratingscombinedget) | **Get** /movie/{movieId}/ratingscombined | Get RT and IMDB movie ratings combined
*MoviesAPI* | [**MovieMovieIdRecommendationsGet**](docs/MoviesAPI.md#moviemovieidrecommendationsget) | **Get** /movie/{movieId}/recommendations | Get recommended movies
*MoviesAPI* | [**MovieMovieIdSimilarGet**](docs/MoviesAPI.md#moviemovieidsimilarget) | **Get** /movie/{movieId}/similar | Get similar movies
*OtherAPI* | [**KeywordKeywordIdGet**](docs/OtherAPI.md#keywordkeywordidget) | **Get** /keyword/{keywordId} | Get keyword
*OtherAPI* | [**WatchprovidersMoviesGet**](docs/OtherAPI.md#watchprovidersmoviesget) | **Get** /watchproviders/movies | Get watch provider movies
*OtherAPI* | [**WatchprovidersRegionsGet**](docs/OtherAPI.md#watchprovidersregionsget) | **Get** /watchproviders/regions | Get watch provider regions
*OtherAPI* | [**WatchprovidersTvGet**](docs/OtherAPI.md#watchproviderstvget) | **Get** /watchproviders/tv | Get watch provider series
*PersonAPI* | [**PersonPersonIdCombinedCreditsGet**](docs/PersonAPI.md#personpersonidcombinedcreditsget) | **Get** /person/{personId}/combined_credits | Get combined credits
*PersonAPI* | [**PersonPersonIdGet**](docs/PersonAPI.md#personpersonidget) | **Get** /person/{personId} | Get person details
*PublicAPI* | [**StatusAppdataGet**](docs/PublicAPI.md#statusappdataget) | **Get** /status/appdata | Get application data volume status
*PublicAPI* | [**StatusGet**](docs/PublicAPI.md#statusget) | **Get** /status | Get Overseerr status
*RequestAPI* | [**RequestCountGet**](docs/RequestAPI.md#requestcountget) | **Get** /request/count | Gets request counts
*RequestAPI* | [**RequestGet**](docs/RequestAPI.md#requestget) | **Get** /request | Get all requests
*RequestAPI* | [**RequestPost**](docs/RequestAPI.md#requestpost) | **Post** /request | Create new request
*RequestAPI* | [**RequestRequestIdDelete**](docs/RequestAPI.md#requestrequestiddelete) | **Delete** /request/{requestId} | Delete request
*RequestAPI* | [**RequestRequestIdGet**](docs/RequestAPI.md#requestrequestidget) | **Get** /request/{requestId} | Get MediaRequest
*RequestAPI* | [**RequestRequestIdPut**](docs/RequestAPI.md#requestrequestidput) | **Put** /request/{requestId} | Update MediaRequest
*RequestAPI* | [**RequestRequestIdRetryPost**](docs/RequestAPI.md#requestrequestidretrypost) | **Post** /request/{requestId}/retry | Retry failed request
*RequestAPI* | [**RequestRequestIdStatusPost**](docs/RequestAPI.md#requestrequestidstatuspost) | **Post** /request/{requestId}/{status} | Update a request&#39;s status
*SearchAPI* | [**DiscoverGenresliderMovieGet**](docs/SearchAPI.md#discovergenreslidermovieget) | **Get** /discover/genreslider/movie | Get genre slider data for movies
*SearchAPI* | [**DiscoverGenresliderTvGet**](docs/SearchAPI.md#discovergenreslidertvget) | **Get** /discover/genreslider/tv | Get genre slider data for TV series
*SearchAPI* | [**DiscoverKeywordKeywordIdMoviesGet**](docs/SearchAPI.md#discoverkeywordkeywordidmoviesget) | **Get** /discover/keyword/{keywordId}/movies | Get movies from keyword
*SearchAPI* | [**DiscoverMoviesGenreGenreIdGet**](docs/SearchAPI.md#discovermoviesgenregenreidget) | **Get** /discover/movies/genre/{genreId} | Discover movies by genre
*SearchAPI* | [**DiscoverMoviesGet**](docs/SearchAPI.md#discovermoviesget) | **Get** /discover/movies | Discover movies
*SearchAPI* | [**DiscoverMoviesLanguageLanguageGet**](docs/SearchAPI.md#discovermovieslanguagelanguageget) | **Get** /discover/movies/language/{language} | Discover movies by original language
*SearchAPI* | [**DiscoverMoviesStudioStudioIdGet**](docs/SearchAPI.md#discovermoviesstudiostudioidget) | **Get** /discover/movies/studio/{studioId} | Discover movies by studio
*SearchAPI* | [**DiscoverMoviesUpcomingGet**](docs/SearchAPI.md#discovermoviesupcomingget) | **Get** /discover/movies/upcoming | Upcoming movies
*SearchAPI* | [**DiscoverTrendingGet**](docs/SearchAPI.md#discovertrendingget) | **Get** /discover/trending | Trending movies and TV
*SearchAPI* | [**DiscoverTvGenreGenreIdGet**](docs/SearchAPI.md#discovertvgenregenreidget) | **Get** /discover/tv/genre/{genreId} | Discover TV shows by genre
*SearchAPI* | [**DiscoverTvGet**](docs/SearchAPI.md#discovertvget) | **Get** /discover/tv | Discover TV shows
*SearchAPI* | [**DiscoverTvLanguageLanguageGet**](docs/SearchAPI.md#discovertvlanguagelanguageget) | **Get** /discover/tv/language/{language} | Discover TV shows by original language
*SearchAPI* | [**DiscoverTvNetworkNetworkIdGet**](docs/SearchAPI.md#discovertvnetworknetworkidget) | **Get** /discover/tv/network/{networkId} | Discover TV shows by network
*SearchAPI* | [**DiscoverTvUpcomingGet**](docs/SearchAPI.md#discovertvupcomingget) | **Get** /discover/tv/upcoming | Discover Upcoming TV shows
*SearchAPI* | [**DiscoverWatchlistGet**](docs/SearchAPI.md#discoverwatchlistget) | **Get** /discover/watchlist | Get the Plex watchlist.
*SearchAPI* | [**SearchCompanyGet**](docs/SearchAPI.md#searchcompanyget) | **Get** /search/company | Search for companies
*SearchAPI* | [**SearchGet**](docs/SearchAPI.md#searchget) | **Get** /search | Search for movies, TV shows, or people
*SearchAPI* | [**SearchKeywordGet**](docs/SearchAPI.md#searchkeywordget) | **Get** /search/keyword | Search for keywords
*ServiceAPI* | [**ServiceRadarrGet**](docs/ServiceAPI.md#serviceradarrget) | **Get** /service/radarr | Get non-sensitive Radarr server list
*ServiceAPI* | [**ServiceRadarrRadarrIdGet**](docs/ServiceAPI.md#serviceradarrradarridget) | **Get** /service/radarr/{radarrId} | Get Radarr server quality profiles and root folders
*ServiceAPI* | [**ServiceSonarrGet**](docs/ServiceAPI.md#servicesonarrget) | **Get** /service/sonarr | Get non-sensitive Sonarr server list
*ServiceAPI* | [**ServiceSonarrLookupTmdbIdGet**](docs/ServiceAPI.md#servicesonarrlookuptmdbidget) | **Get** /service/sonarr/lookup/{tmdbId} | Get series from Sonarr
*ServiceAPI* | [**ServiceSonarrSonarrIdGet**](docs/ServiceAPI.md#servicesonarrsonarridget) | **Get** /service/sonarr/{sonarrId} | Get Sonarr server quality profiles and root folders
*SettingsAPI* | [**SettingsAboutGet**](docs/SettingsAPI.md#settingsaboutget) | **Get** /settings/about | Get server stats
*SettingsAPI* | [**SettingsCacheCacheIdFlushPost**](docs/SettingsAPI.md#settingscachecacheidflushpost) | **Post** /settings/cache/{cacheId}/flush | Flush a specific cache
*SettingsAPI* | [**SettingsCacheGet**](docs/SettingsAPI.md#settingscacheget) | **Get** /settings/cache | Get a list of active caches
*SettingsAPI* | [**SettingsDiscoverAddPost**](docs/SettingsAPI.md#settingsdiscoveraddpost) | **Post** /settings/discover/add | Add a new slider
*SettingsAPI* | [**SettingsDiscoverGet**](docs/SettingsAPI.md#settingsdiscoverget) | **Get** /settings/discover | Get all discover sliders
*SettingsAPI* | [**SettingsDiscoverPost**](docs/SettingsAPI.md#settingsdiscoverpost) | **Post** /settings/discover | Batch update all sliders.
*SettingsAPI* | [**SettingsDiscoverResetGet**](docs/SettingsAPI.md#settingsdiscoverresetget) | **Get** /settings/discover/reset | Reset all discover sliders
*SettingsAPI* | [**SettingsDiscoverSliderIdDelete**](docs/SettingsAPI.md#settingsdiscoverslideriddelete) | **Delete** /settings/discover/{sliderId} | Delete slider by ID
*SettingsAPI* | [**SettingsDiscoverSliderIdPut**](docs/SettingsAPI.md#settingsdiscoverslideridput) | **Put** /settings/discover/{sliderId} | Update a single slider
*SettingsAPI* | [**SettingsInitializePost**](docs/SettingsAPI.md#settingsinitializepost) | **Post** /settings/initialize | Initialize application
*SettingsAPI* | [**SettingsJobsGet**](docs/SettingsAPI.md#settingsjobsget) | **Get** /settings/jobs | Get scheduled jobs
*SettingsAPI* | [**SettingsJobsJobIdCancelPost**](docs/SettingsAPI.md#settingsjobsjobidcancelpost) | **Post** /settings/jobs/{jobId}/cancel | Cancel a specific job
*SettingsAPI* | [**SettingsJobsJobIdRunPost**](docs/SettingsAPI.md#settingsjobsjobidrunpost) | **Post** /settings/jobs/{jobId}/run | Invoke a specific job
*SettingsAPI* | [**SettingsJobsJobIdSchedulePost**](docs/SettingsAPI.md#settingsjobsjobidschedulepost) | **Post** /settings/jobs/{jobId}/schedule | Modify job schedule
*SettingsAPI* | [**SettingsLogsGet**](docs/SettingsAPI.md#settingslogsget) | **Get** /settings/logs | Returns logs
*SettingsAPI* | [**SettingsMainGet**](docs/SettingsAPI.md#settingsmainget) | **Get** /settings/main | Get main settings
*SettingsAPI* | [**SettingsMainPost**](docs/SettingsAPI.md#settingsmainpost) | **Post** /settings/main | Update main settings
*SettingsAPI* | [**SettingsMainRegeneratePost**](docs/SettingsAPI.md#settingsmainregeneratepost) | **Post** /settings/main/regenerate | Get main settings with newly-generated API key
*SettingsAPI* | [**SettingsNotificationsDiscordGet**](docs/SettingsAPI.md#settingsnotificationsdiscordget) | **Get** /settings/notifications/discord | Get Discord notification settings
*SettingsAPI* | [**SettingsNotificationsDiscordPost**](docs/SettingsAPI.md#settingsnotificationsdiscordpost) | **Post** /settings/notifications/discord | Update Discord notification settings
*SettingsAPI* | [**SettingsNotificationsDiscordTestPost**](docs/SettingsAPI.md#settingsnotificationsdiscordtestpost) | **Post** /settings/notifications/discord/test | Test Discord settings
*SettingsAPI* | [**SettingsNotificationsEmailGet**](docs/SettingsAPI.md#settingsnotificationsemailget) | **Get** /settings/notifications/email | Get email notification settings
*SettingsAPI* | [**SettingsNotificationsEmailPost**](docs/SettingsAPI.md#settingsnotificationsemailpost) | **Post** /settings/notifications/email | Update email notification settings
*SettingsAPI* | [**SettingsNotificationsEmailTestPost**](docs/SettingsAPI.md#settingsnotificationsemailtestpost) | **Post** /settings/notifications/email/test | Test email settings
*SettingsAPI* | [**SettingsNotificationsGotifyGet**](docs/SettingsAPI.md#settingsnotificationsgotifyget) | **Get** /settings/notifications/gotify | Get Gotify notification settings
*SettingsAPI* | [**SettingsNotificationsGotifyPost**](docs/SettingsAPI.md#settingsnotificationsgotifypost) | **Post** /settings/notifications/gotify | Update Gotify notification settings
*SettingsAPI* | [**SettingsNotificationsGotifyTestPost**](docs/SettingsAPI.md#settingsnotificationsgotifytestpost) | **Post** /settings/notifications/gotify/test | Test Gotify settings
*SettingsAPI* | [**SettingsNotificationsLunaseaGet**](docs/SettingsAPI.md#settingsnotificationslunaseaget) | **Get** /settings/notifications/lunasea | Get LunaSea notification settings
*SettingsAPI* | [**SettingsNotificationsLunaseaPost**](docs/SettingsAPI.md#settingsnotificationslunaseapost) | **Post** /settings/notifications/lunasea | Update LunaSea notification settings
*SettingsAPI* | [**SettingsNotificationsLunaseaTestPost**](docs/SettingsAPI.md#settingsnotificationslunaseatestpost) | **Post** /settings/notifications/lunasea/test | Test LunaSea settings
*SettingsAPI* | [**SettingsNotificationsPushbulletGet**](docs/SettingsAPI.md#settingsnotificationspushbulletget) | **Get** /settings/notifications/pushbullet | Get Pushbullet notification settings
*SettingsAPI* | [**SettingsNotificationsPushbulletPost**](docs/SettingsAPI.md#settingsnotificationspushbulletpost) | **Post** /settings/notifications/pushbullet | Update Pushbullet notification settings
*SettingsAPI* | [**SettingsNotificationsPushbulletTestPost**](docs/SettingsAPI.md#settingsnotificationspushbullettestpost) | **Post** /settings/notifications/pushbullet/test | Test Pushbullet settings
*SettingsAPI* | [**SettingsNotificationsPushoverGet**](docs/SettingsAPI.md#settingsnotificationspushoverget) | **Get** /settings/notifications/pushover | Get Pushover notification settings
*SettingsAPI* | [**SettingsNotificationsPushoverPost**](docs/SettingsAPI.md#settingsnotificationspushoverpost) | **Post** /settings/notifications/pushover | Update Pushover notification settings
*SettingsAPI* | [**SettingsNotificationsPushoverSoundsGet**](docs/SettingsAPI.md#settingsnotificationspushoversoundsget) | **Get** /settings/notifications/pushover/sounds | Get Pushover sounds
*SettingsAPI* | [**SettingsNotificationsPushoverTestPost**](docs/SettingsAPI.md#settingsnotificationspushovertestpost) | **Post** /settings/notifications/pushover/test | Test Pushover settings
*SettingsAPI* | [**SettingsNotificationsSlackGet**](docs/SettingsAPI.md#settingsnotificationsslackget) | **Get** /settings/notifications/slack | Get Slack notification settings
*SettingsAPI* | [**SettingsNotificationsSlackPost**](docs/SettingsAPI.md#settingsnotificationsslackpost) | **Post** /settings/notifications/slack | Update Slack notification settings
*SettingsAPI* | [**SettingsNotificationsSlackTestPost**](docs/SettingsAPI.md#settingsnotificationsslacktestpost) | **Post** /settings/notifications/slack/test | Test Slack settings
*SettingsAPI* | [**SettingsNotificationsTelegramGet**](docs/SettingsAPI.md#settingsnotificationstelegramget) | **Get** /settings/notifications/telegram | Get Telegram notification settings
*SettingsAPI* | [**SettingsNotificationsTelegramPost**](docs/SettingsAPI.md#settingsnotificationstelegrampost) | **Post** /settings/notifications/telegram | Update Telegram notification settings
*SettingsAPI* | [**SettingsNotificationsTelegramTestPost**](docs/SettingsAPI.md#settingsnotificationstelegramtestpost) | **Post** /settings/notifications/telegram/test | Test Telegram settings
*SettingsAPI* | [**SettingsNotificationsWebhookGet**](docs/SettingsAPI.md#settingsnotificationswebhookget) | **Get** /settings/notifications/webhook | Get webhook notification settings
*SettingsAPI* | [**SettingsNotificationsWebhookPost**](docs/SettingsAPI.md#settingsnotificationswebhookpost) | **Post** /settings/notifications/webhook | Update webhook notification settings
*SettingsAPI* | [**SettingsNotificationsWebhookTestPost**](docs/SettingsAPI.md#settingsnotificationswebhooktestpost) | **Post** /settings/notifications/webhook/test | Test webhook settings
*SettingsAPI* | [**SettingsNotificationsWebpushGet**](docs/SettingsAPI.md#settingsnotificationswebpushget) | **Get** /settings/notifications/webpush | Get Web Push notification settings
*SettingsAPI* | [**SettingsNotificationsWebpushPost**](docs/SettingsAPI.md#settingsnotificationswebpushpost) | **Post** /settings/notifications/webpush | Update Web Push notification settings
*SettingsAPI* | [**SettingsNotificationsWebpushTestPost**](docs/SettingsAPI.md#settingsnotificationswebpushtestpost) | **Post** /settings/notifications/webpush/test | Test Web Push settings
*SettingsAPI* | [**SettingsPlexDevicesServersGet**](docs/SettingsAPI.md#settingsplexdevicesserversget) | **Get** /settings/plex/devices/servers | Gets the user&#39;s available Plex servers
*SettingsAPI* | [**SettingsPlexGet**](docs/SettingsAPI.md#settingsplexget) | **Get** /settings/plex | Get Plex settings
*SettingsAPI* | [**SettingsPlexLibraryGet**](docs/SettingsAPI.md#settingsplexlibraryget) | **Get** /settings/plex/library | Get Plex libraries
*SettingsAPI* | [**SettingsPlexPost**](docs/SettingsAPI.md#settingsplexpost) | **Post** /settings/plex | Update Plex settings
*SettingsAPI* | [**SettingsPlexSyncGet**](docs/SettingsAPI.md#settingsplexsyncget) | **Get** /settings/plex/sync | Get status of full Plex library scan
*SettingsAPI* | [**SettingsPlexSyncPost**](docs/SettingsAPI.md#settingsplexsyncpost) | **Post** /settings/plex/sync | Start full Plex library scan
*SettingsAPI* | [**SettingsPlexUsersGet**](docs/SettingsAPI.md#settingsplexusersget) | **Get** /settings/plex/users | Get Plex users
*SettingsAPI* | [**SettingsPublicGet**](docs/SettingsAPI.md#settingspublicget) | **Get** /settings/public | Get public settings
*SettingsAPI* | [**SettingsRadarrGet**](docs/SettingsAPI.md#settingsradarrget) | **Get** /settings/radarr | Get Radarr settings
*SettingsAPI* | [**SettingsRadarrPost**](docs/SettingsAPI.md#settingsradarrpost) | **Post** /settings/radarr | Create Radarr instance
*SettingsAPI* | [**SettingsRadarrRadarrIdDelete**](docs/SettingsAPI.md#settingsradarrradarriddelete) | **Delete** /settings/radarr/{radarrId} | Delete Radarr instance
*SettingsAPI* | [**SettingsRadarrRadarrIdProfilesGet**](docs/SettingsAPI.md#settingsradarrradarridprofilesget) | **Get** /settings/radarr/{radarrId}/profiles | Get available Radarr profiles
*SettingsAPI* | [**SettingsRadarrRadarrIdPut**](docs/SettingsAPI.md#settingsradarrradarridput) | **Put** /settings/radarr/{radarrId} | Update Radarr instance
*SettingsAPI* | [**SettingsRadarrTestPost**](docs/SettingsAPI.md#settingsradarrtestpost) | **Post** /settings/radarr/test | Test Radarr configuration
*SettingsAPI* | [**SettingsSonarrGet**](docs/SettingsAPI.md#settingssonarrget) | **Get** /settings/sonarr | Get Sonarr settings
*SettingsAPI* | [**SettingsSonarrPost**](docs/SettingsAPI.md#settingssonarrpost) | **Post** /settings/sonarr | Create Sonarr instance
*SettingsAPI* | [**SettingsSonarrSonarrIdDelete**](docs/SettingsAPI.md#settingssonarrsonarriddelete) | **Delete** /settings/sonarr/{sonarrId} | Delete Sonarr instance
*SettingsAPI* | [**SettingsSonarrSonarrIdPut**](docs/SettingsAPI.md#settingssonarrsonarridput) | **Put** /settings/sonarr/{sonarrId} | Update Sonarr instance
*SettingsAPI* | [**SettingsSonarrTestPost**](docs/SettingsAPI.md#settingssonarrtestpost) | **Post** /settings/sonarr/test | Test Sonarr configuration
*SettingsAPI* | [**SettingsTautulliGet**](docs/SettingsAPI.md#settingstautulliget) | **Get** /settings/tautulli | Get Tautulli settings
*SettingsAPI* | [**SettingsTautulliPost**](docs/SettingsAPI.md#settingstautullipost) | **Post** /settings/tautulli | Update Tautulli settings
*TmdbAPI* | [**BackdropsGet**](docs/TmdbAPI.md#backdropsget) | **Get** /backdrops | Get backdrops of trending items
*TmdbAPI* | [**GenresMovieGet**](docs/TmdbAPI.md#genresmovieget) | **Get** /genres/movie | Get list of official TMDB movie genres
*TmdbAPI* | [**GenresTvGet**](docs/TmdbAPI.md#genrestvget) | **Get** /genres/tv | Get list of official TMDB movie genres
*TmdbAPI* | [**LanguagesGet**](docs/TmdbAPI.md#languagesget) | **Get** /languages | Languages supported by TMDB
*TmdbAPI* | [**NetworkNetworkIdGet**](docs/TmdbAPI.md#networknetworkidget) | **Get** /network/{networkId} | Get TV network details
*TmdbAPI* | [**RegionsGet**](docs/TmdbAPI.md#regionsget) | **Get** /regions | Regions supported by TMDB
*TmdbAPI* | [**StudioStudioIdGet**](docs/TmdbAPI.md#studiostudioidget) | **Get** /studio/{studioId} | Get movie studio details
*TvAPI* | [**TvTvIdGet**](docs/TvAPI.md#tvtvidget) | **Get** /tv/{tvId} | Get TV details
*TvAPI* | [**TvTvIdRatingsGet**](docs/TvAPI.md#tvtvidratingsget) | **Get** /tv/{tvId}/ratings | Get TV ratings
*TvAPI* | [**TvTvIdRecommendationsGet**](docs/TvAPI.md#tvtvidrecommendationsget) | **Get** /tv/{tvId}/recommendations | Get recommended TV series
*TvAPI* | [**TvTvIdSeasonSeasonIdGet**](docs/TvAPI.md#tvtvidseasonseasonidget) | **Get** /tv/{tvId}/season/{seasonId} | Get season details and episode list
*TvAPI* | [**TvTvIdSimilarGet**](docs/TvAPI.md#tvtvidsimilarget) | **Get** /tv/{tvId}/similar | Get similar TV series
*UsersAPI* | [**AuthMeGet**](docs/UsersAPI.md#authmeget) | **Get** /auth/me | Get logged-in user
*UsersAPI* | [**AuthResetPasswordGuidPost**](docs/UsersAPI.md#authresetpasswordguidpost) | **Post** /auth/reset-password/{guid} | Reset the password for a user
*UsersAPI* | [**AuthResetPasswordPost**](docs/UsersAPI.md#authresetpasswordpost) | **Post** /auth/reset-password | Send a reset password email
*UsersAPI* | [**SettingsPlexUsersGet**](docs/UsersAPI.md#settingsplexusersget) | **Get** /settings/plex/users | Get Plex users
*UsersAPI* | [**UserGet**](docs/UsersAPI.md#userget) | **Get** /user | Get all users
*UsersAPI* | [**UserImportFromPlexPost**](docs/UsersAPI.md#userimportfromplexpost) | **Post** /user/import-from-plex | Import all users from Plex
*UsersAPI* | [**UserPost**](docs/UsersAPI.md#userpost) | **Post** /user | Create new user
*UsersAPI* | [**UserPut**](docs/UsersAPI.md#userput) | **Put** /user | Update batch of users
*UsersAPI* | [**UserRegisterPushSubscriptionPost**](docs/UsersAPI.md#userregisterpushsubscriptionpost) | **Post** /user/registerPushSubscription | Register a web push /user/registerPushSubscription
*UsersAPI* | [**UserUserIdDelete**](docs/UsersAPI.md#useruseriddelete) | **Delete** /user/{userId} | Delete user by ID
*UsersAPI* | [**UserUserIdGet**](docs/UsersAPI.md#useruseridget) | **Get** /user/{userId} | Get user by ID
*UsersAPI* | [**UserUserIdPut**](docs/UsersAPI.md#useruseridput) | **Put** /user/{userId} | Update a user by user ID
*UsersAPI* | [**UserUserIdQuotaGet**](docs/UsersAPI.md#useruseridquotaget) | **Get** /user/{userId}/quota | Get quotas for a specific user
*UsersAPI* | [**UserUserIdRequestsGet**](docs/UsersAPI.md#useruseridrequestsget) | **Get** /user/{userId}/requests | Get requests for a specific user
*UsersAPI* | [**UserUserIdSettingsMainGet**](docs/UsersAPI.md#useruseridsettingsmainget) | **Get** /user/{userId}/settings/main | Get general settings for a user
*UsersAPI* | [**UserUserIdSettingsMainPost**](docs/UsersAPI.md#useruseridsettingsmainpost) | **Post** /user/{userId}/settings/main | Update general settings for a user
*UsersAPI* | [**UserUserIdSettingsNotificationsGet**](docs/UsersAPI.md#useruseridsettingsnotificationsget) | **Get** /user/{userId}/settings/notifications | Get notification settings for a user
*UsersAPI* | [**UserUserIdSettingsNotificationsPost**](docs/UsersAPI.md#useruseridsettingsnotificationspost) | **Post** /user/{userId}/settings/notifications | Update notification settings for a user
*UsersAPI* | [**UserUserIdSettingsPasswordGet**](docs/UsersAPI.md#useruseridsettingspasswordget) | **Get** /user/{userId}/settings/password | Get password page informatiom
*UsersAPI* | [**UserUserIdSettingsPasswordPost**](docs/UsersAPI.md#useruseridsettingspasswordpost) | **Post** /user/{userId}/settings/password | Update password for a user
*UsersAPI* | [**UserUserIdSettingsPermissionsGet**](docs/UsersAPI.md#useruseridsettingspermissionsget) | **Get** /user/{userId}/settings/permissions | Get permission settings for a user
*UsersAPI* | [**UserUserIdSettingsPermissionsPost**](docs/UsersAPI.md#useruseridsettingspermissionspost) | **Post** /user/{userId}/settings/permissions | Update permission settings for a user
*UsersAPI* | [**UserUserIdWatchDataGet**](docs/UsersAPI.md#useruseridwatchdataget) | **Get** /user/{userId}/watch_data | Get watch data
*UsersAPI* | [**UserUserIdWatchlistGet**](docs/UsersAPI.md#useruseridwatchlistget) | **Get** /user/{userId}/watchlist | Get the Plex watchlist for a specific user


## Documentation For Models

 - [AuthLocalPostRequest](docs/AuthLocalPostRequest.md)
 - [AuthLogoutPost200Response](docs/AuthLogoutPost200Response.md)
 - [AuthPlexPostRequest](docs/AuthPlexPostRequest.md)
 - [AuthResetPasswordGuidPostRequest](docs/AuthResetPasswordGuidPostRequest.md)
 - [AuthResetPasswordPostRequest](docs/AuthResetPasswordPostRequest.md)
 - [Cast](docs/Cast.md)
 - [Collection](docs/Collection.md)
 - [Company](docs/Company.md)
 - [CreditCast](docs/CreditCast.md)
 - [CreditCrew](docs/CreditCrew.md)
 - [Crew](docs/Crew.md)
 - [DiscordSettings](docs/DiscordSettings.md)
 - [DiscordSettingsOptions](docs/DiscordSettingsOptions.md)
 - [DiscoverGenresliderMovieGet200ResponseInner](docs/DiscoverGenresliderMovieGet200ResponseInner.md)
 - [DiscoverMoviesGenreGenreIdGet200Response](docs/DiscoverMoviesGenreGenreIdGet200Response.md)
 - [DiscoverMoviesGet200Response](docs/DiscoverMoviesGet200Response.md)
 - [DiscoverMoviesLanguageLanguageGet200Response](docs/DiscoverMoviesLanguageLanguageGet200Response.md)
 - [DiscoverMoviesStudioStudioIdGet200Response](docs/DiscoverMoviesStudioStudioIdGet200Response.md)
 - [DiscoverSlider](docs/DiscoverSlider.md)
 - [DiscoverTvGenreGenreIdGet200Response](docs/DiscoverTvGenreGenreIdGet200Response.md)
 - [DiscoverTvGet200Response](docs/DiscoverTvGet200Response.md)
 - [DiscoverTvLanguageLanguageGet200Response](docs/DiscoverTvLanguageLanguageGet200Response.md)
 - [DiscoverTvNetworkNetworkIdGet200Response](docs/DiscoverTvNetworkNetworkIdGet200Response.md)
 - [Episode](docs/Episode.md)
 - [ExternalIds](docs/ExternalIds.md)
 - [Genre](docs/Genre.md)
 - [GenresMovieGet200ResponseInner](docs/GenresMovieGet200ResponseInner.md)
 - [GenresTvGet200ResponseInner](docs/GenresTvGet200ResponseInner.md)
 - [GotifySettings](docs/GotifySettings.md)
 - [GotifySettingsOptions](docs/GotifySettingsOptions.md)
 - [Issue](docs/Issue.md)
 - [IssueComment](docs/IssueComment.md)
 - [IssueCommentCommentIdPutRequest](docs/IssueCommentCommentIdPutRequest.md)
 - [IssueCountGet200Response](docs/IssueCountGet200Response.md)
 - [IssueGet200Response](docs/IssueGet200Response.md)
 - [IssueIssueIdCommentPostRequest](docs/IssueIssueIdCommentPostRequest.md)
 - [IssuePostRequest](docs/IssuePostRequest.md)
 - [Job](docs/Job.md)
 - [Keyword](docs/Keyword.md)
 - [LanguagesGet200ResponseInner](docs/LanguagesGet200ResponseInner.md)
 - [LunaSeaSettings](docs/LunaSeaSettings.md)
 - [LunaSeaSettingsOptions](docs/LunaSeaSettingsOptions.md)
 - [MainSettings](docs/MainSettings.md)
 - [MediaGet200Response](docs/MediaGet200Response.md)
 - [MediaInfo](docs/MediaInfo.md)
 - [MediaMediaIdStatusPostRequest](docs/MediaMediaIdStatusPostRequest.md)
 - [MediaMediaIdWatchDataGet200Response](docs/MediaMediaIdWatchDataGet200Response.md)
 - [MediaMediaIdWatchDataGet200ResponseData](docs/MediaMediaIdWatchDataGet200ResponseData.md)
 - [MediaRequest](docs/MediaRequest.md)
 - [MediaRequestModifiedBy](docs/MediaRequestModifiedBy.md)
 - [MovieDetails](docs/MovieDetails.md)
 - [MovieDetailsCollection](docs/MovieDetailsCollection.md)
 - [MovieDetailsCredits](docs/MovieDetailsCredits.md)
 - [MovieDetailsProductionCountriesInner](docs/MovieDetailsProductionCountriesInner.md)
 - [MovieDetailsReleases](docs/MovieDetailsReleases.md)
 - [MovieDetailsReleasesResultsInner](docs/MovieDetailsReleasesResultsInner.md)
 - [MovieDetailsReleasesResultsInnerReleaseDatesInner](docs/MovieDetailsReleasesResultsInnerReleaseDatesInner.md)
 - [MovieMovieIdRatingsGet200Response](docs/MovieMovieIdRatingsGet200Response.md)
 - [MovieMovieIdRatingscombinedGet200Response](docs/MovieMovieIdRatingscombinedGet200Response.md)
 - [MovieMovieIdRatingscombinedGet200ResponseImdb](docs/MovieMovieIdRatingscombinedGet200ResponseImdb.md)
 - [MovieResult](docs/MovieResult.md)
 - [Network](docs/Network.md)
 - [NotificationAgentTypes](docs/NotificationAgentTypes.md)
 - [NotificationEmailSettings](docs/NotificationEmailSettings.md)
 - [NotificationEmailSettingsOptions](docs/NotificationEmailSettingsOptions.md)
 - [PageInfo](docs/PageInfo.md)
 - [PersonDetails](docs/PersonDetails.md)
 - [PersonPersonIdCombinedCreditsGet200Response](docs/PersonPersonIdCombinedCreditsGet200Response.md)
 - [PersonResult](docs/PersonResult.md)
 - [PersonResultKnownForInner](docs/PersonResultKnownForInner.md)
 - [PlexConnection](docs/PlexConnection.md)
 - [PlexDevice](docs/PlexDevice.md)
 - [PlexLibrary](docs/PlexLibrary.md)
 - [PlexSettings](docs/PlexSettings.md)
 - [ProductionCompany](docs/ProductionCompany.md)
 - [PublicSettings](docs/PublicSettings.md)
 - [PushbulletSettings](docs/PushbulletSettings.md)
 - [PushbulletSettingsOptions](docs/PushbulletSettingsOptions.md)
 - [PushoverSettings](docs/PushoverSettings.md)
 - [PushoverSettingsOptions](docs/PushoverSettingsOptions.md)
 - [RadarrSettings](docs/RadarrSettings.md)
 - [RegionsGet200ResponseInner](docs/RegionsGet200ResponseInner.md)
 - [RelatedVideo](docs/RelatedVideo.md)
 - [RequestCountGet200Response](docs/RequestCountGet200Response.md)
 - [RequestPostRequest](docs/RequestPostRequest.md)
 - [RequestPostRequestSeasons](docs/RequestPostRequestSeasons.md)
 - [RequestRequestIdPutRequest](docs/RequestRequestIdPutRequest.md)
 - [SearchCompanyGet200Response](docs/SearchCompanyGet200Response.md)
 - [SearchGet200Response](docs/SearchGet200Response.md)
 - [SearchGet200ResponseResultsInner](docs/SearchGet200ResponseResultsInner.md)
 - [SearchKeywordGet200Response](docs/SearchKeywordGet200Response.md)
 - [Season](docs/Season.md)
 - [ServarrTag](docs/ServarrTag.md)
 - [ServiceProfile](docs/ServiceProfile.md)
 - [ServiceRadarrRadarrIdGet200Response](docs/ServiceRadarrRadarrIdGet200Response.md)
 - [ServiceSonarrSonarrIdGet200Response](docs/ServiceSonarrSonarrIdGet200Response.md)
 - [SettingsAboutGet200Response](docs/SettingsAboutGet200Response.md)
 - [SettingsCacheGet200Response](docs/SettingsCacheGet200Response.md)
 - [SettingsCacheGet200ResponseApiCachesInner](docs/SettingsCacheGet200ResponseApiCachesInner.md)
 - [SettingsCacheGet200ResponseApiCachesInnerStats](docs/SettingsCacheGet200ResponseApiCachesInnerStats.md)
 - [SettingsCacheGet200ResponseImageCache](docs/SettingsCacheGet200ResponseImageCache.md)
 - [SettingsCacheGet200ResponseImageCacheTmdb](docs/SettingsCacheGet200ResponseImageCacheTmdb.md)
 - [SettingsDiscoverAddPostRequest](docs/SettingsDiscoverAddPostRequest.md)
 - [SettingsDiscoverSliderIdPutRequest](docs/SettingsDiscoverSliderIdPutRequest.md)
 - [SettingsJobsJobIdSchedulePostRequest](docs/SettingsJobsJobIdSchedulePostRequest.md)
 - [SettingsLogsGet200ResponseInner](docs/SettingsLogsGet200ResponseInner.md)
 - [SettingsNotificationsPushoverSoundsGet200ResponseInner](docs/SettingsNotificationsPushoverSoundsGet200ResponseInner.md)
 - [SettingsPlexSyncGet200Response](docs/SettingsPlexSyncGet200Response.md)
 - [SettingsPlexSyncPostRequest](docs/SettingsPlexSyncPostRequest.md)
 - [SettingsPlexUsersGet200ResponseInner](docs/SettingsPlexUsersGet200ResponseInner.md)
 - [SettingsRadarrTestPost200Response](docs/SettingsRadarrTestPost200Response.md)
 - [SettingsRadarrTestPostRequest](docs/SettingsRadarrTestPostRequest.md)
 - [SettingsSonarrTestPostRequest](docs/SettingsSonarrTestPostRequest.md)
 - [SlackSettings](docs/SlackSettings.md)
 - [SlackSettingsOptions](docs/SlackSettingsOptions.md)
 - [SonarrSeries](docs/SonarrSeries.md)
 - [SonarrSeriesAddOptionsInner](docs/SonarrSeriesAddOptionsInner.md)
 - [SonarrSeriesImagesInner](docs/SonarrSeriesImagesInner.md)
 - [SonarrSeriesRatingsInner](docs/SonarrSeriesRatingsInner.md)
 - [SonarrSeriesSeasonsInner](docs/SonarrSeriesSeasonsInner.md)
 - [SonarrSettings](docs/SonarrSettings.md)
 - [SpokenLanguage](docs/SpokenLanguage.md)
 - [StatusAppdataGet200Response](docs/StatusAppdataGet200Response.md)
 - [StatusGet200Response](docs/StatusGet200Response.md)
 - [TautulliSettings](docs/TautulliSettings.md)
 - [TelegramSettings](docs/TelegramSettings.md)
 - [TelegramSettingsOptions](docs/TelegramSettingsOptions.md)
 - [TvDetails](docs/TvDetails.md)
 - [TvDetailsContentRatings](docs/TvDetailsContentRatings.md)
 - [TvDetailsContentRatingsResultsInner](docs/TvDetailsContentRatingsResultsInner.md)
 - [TvDetailsCreatedByInner](docs/TvDetailsCreatedByInner.md)
 - [TvResult](docs/TvResult.md)
 - [TvTvIdRatingsGet200Response](docs/TvTvIdRatingsGet200Response.md)
 - [User](docs/User.md)
 - [UserGet200Response](docs/UserGet200Response.md)
 - [UserImportFromPlexPostRequest](docs/UserImportFromPlexPostRequest.md)
 - [UserPostRequest](docs/UserPostRequest.md)
 - [UserPutRequest](docs/UserPutRequest.md)
 - [UserRegisterPushSubscriptionPostRequest](docs/UserRegisterPushSubscriptionPostRequest.md)
 - [UserSettings](docs/UserSettings.md)
 - [UserSettingsNotifications](docs/UserSettingsNotifications.md)
 - [UserUserIdQuotaGet200Response](docs/UserUserIdQuotaGet200Response.md)
 - [UserUserIdQuotaGet200ResponseMovie](docs/UserUserIdQuotaGet200ResponseMovie.md)
 - [UserUserIdRequestsGet200Response](docs/UserUserIdRequestsGet200Response.md)
 - [UserUserIdSettingsMainGet200Response](docs/UserUserIdSettingsMainGet200Response.md)
 - [UserUserIdSettingsMainPostRequest](docs/UserUserIdSettingsMainPostRequest.md)
 - [UserUserIdSettingsPasswordGet200Response](docs/UserUserIdSettingsPasswordGet200Response.md)
 - [UserUserIdSettingsPasswordPostRequest](docs/UserUserIdSettingsPasswordPostRequest.md)
 - [UserUserIdSettingsPermissionsGet200Response](docs/UserUserIdSettingsPermissionsGet200Response.md)
 - [UserUserIdSettingsPermissionsPostRequest](docs/UserUserIdSettingsPermissionsPostRequest.md)
 - [UserUserIdWatchDataGet200Response](docs/UserUserIdWatchDataGet200Response.md)
 - [UserUserIdWatchlistGet200Response](docs/UserUserIdWatchlistGet200Response.md)
 - [UserUserIdWatchlistGet200ResponseResultsInner](docs/UserUserIdWatchlistGet200ResponseResultsInner.md)
 - [WatchProviderDetails](docs/WatchProviderDetails.md)
 - [WatchProviderRegion](docs/WatchProviderRegion.md)
 - [WatchProvidersInner](docs/WatchProvidersInner.md)
 - [WebPushSettings](docs/WebPushSettings.md)
 - [WebhookSettings](docs/WebhookSettings.md)
 - [WebhookSettingsOptions](docs/WebhookSettingsOptions.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### cookieAuth

- **Type**: API key
- **API key parameter name**: connect.sid
- **Location**: 

Note, each API key must be added to a map of `map[string]APIKey` where the key is: connect.sid and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		overseerr_go.ContextAPIKeys,
		map[string]overseerr_go.APIKey{
			"connect.sid": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```

### apiKey

- **Type**: API key
- **API key parameter name**: X-Api-Key
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: X-Api-Key and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		overseerr_go.ContextAPIKeys,
		map[string]overseerr_go.APIKey{
			"X-Api-Key": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



