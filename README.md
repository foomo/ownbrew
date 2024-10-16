# Sesamy CLI

[![Build Status](https://github.com/foomo/sesamy-cli/actions/workflows/test.yml/badge.svg?branch=main&event=push)](https://github.com/foomo/sesamy-cli/actions/workflows/test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/foomo/sesamy-cli)](https://goreportcard.com/report/github.com/foomo/sesamy-cli)
[![godoc](https://godoc.org/github.com/foomo/sesamy-cli?status.svg)](https://godoc.org/github.com/foomo/sesamy-cli)
[![goreleaser](https://github.com/foomo/sesamy-cli/actions/workflows/release.yml/badge.svg)](https://github.com/foomo/sesamy-cli/actions)

> CLI to keep you sane while working with GTM.

## Installing

Install the latest release of the cli:

````bash
$ brew update
$ brew install foomo/tap/sesamy-cli
````

## Usage

Add a `sesamy.yaml` configuration

```yaml
# yaml-language-server: $schema=https://raw.githubusercontent.com/foomo/sesamy-cli/v0.4.1/sesamy.yaml
version: '1.0'

# Whether to redact the visitor ip
redactVisitorIp: true

# --- Google API settings
googleApi:
  # Single line Service Account credentials
  credentials: '{...\\n...\\n...}'
  # Path to the Service Account credentials json file
  credentialsFile: google_service_account_creds.json
  # Current API request quota (send a request to increase the quota)
  requestQuota: 15

# --- Google Tag Manager settings
googleTagManager:
  # The account id
  accountId: 6099238525
  # Web container settings
  webContainer:
    # The container tag id
    tagId: GTM-57BHX34G
    # The container id
    containerId: 175355532
    # The workspace id that should be used by the api
    workspaceId: 23
  # Server container settings
  serverContainer:
    # The container tag id
    tagId: GTM-5NWPR4QW
    # The container id
    containerId: 175348980
    # The workspace id that should be used by the api
    workspaceId: 10

# --- Google Tag settings
googleTag:
  # A tag ID is an identifier that you put on your page to load a given Google tag
  tagId: G-PZ5ELRCR31
  # Whether a page_view should be sent on initial load
  sendPageView: true
  # Enable debug mode for all user devices
  debugMode: false
  # Google Tag Manager web container settings
  webContainer:
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - SelectItem
  # Google Tag Manager server container settings
  serverContainer:
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - SelectItem
  # Google Tag Manager web container settings
  typeScript:
    # Target directory for generate files
    outputPath: path/to/target
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          ## GA4 Automatically collected events
          ## https://support.google.com/analytics/answer/9234069
          - Click
          - FileDownload
          - FirstVisit
          - FormStart
          - FormSubmit
          - PageView
          - Scroll
          - UserEngagement
          - VideoComplete
          - VideoProgress
          - VideoStart
          - ViewSearchResults
          ## Recommended events
          ## https://developers.google.com/tag-platform/gtagjs/reference/events
          - AdImpression
          - AddPaymentInfo
          - AddShippingInfo
          - AddToCart
          - AddToWishlist
          - BeginCheckout
          - CampaignDetails
          - CloseConvertLead
          - CloseUnconvertLead
          - DisqualifyLead
          - EarnVirtualMoney
          - Exception
          - GenerateLead
          - JoinGroup
          - LevelEnd
          - LevelStart
          - LevelUp
          - Login
          - PostScore
          - Purchase
          - QualifyLead
          - Refund
          - RemoveFromCart
          - ScreenView
          - Search
          - SelectContent
          - SelectItem
          - SelectPromotion
          - SessionStart
          - Share
          - SignUp
          - SpendVirtualCurrency
          - TutorialBegin
          - TutorialComplete
          - UnlockAchievement
          - ViewCart
          - ViewItem
          - ViewItemList
          - ViewPromotion
          - WorkingLead


# --- Google Analytics settings
googleAnalytics:
  # Enable provider
  enabled: true
  # Google GTag.js settings
  googleGTag:
    # Provision custom client
    enabled: true
    # Client priority
    priority: 10
    # Patch ecommerce items
    ecommerceItems: true
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: analytics_storage
  # Google Tag Manager web container settings
  webContainer:
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - SelectItem
  # Google Tag Manager server container settings
  serverContainer:
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - SelectItem

# --- Google Ads
googleAds:
  # Enable provider
  enabled: true
  # Google Ads Conversion Tracking ID
  conversionId: ''
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: ad_storage
  # Google Ads Remarketing settings
  remarketing:
    # Enable Google Ads Remarketing
    enabled: true
    # Enable conversion linking
    enableConversionLinker: true
  # Google Ads Conversion settings
  conversion:
    # Enable Google Ads Conversion
    enabled: true
    # Google Ads Conversion Tracking Label
    conversionLabel: ''
    # Google Tag Manager server container settings
    serverContainer:
      # Contemplate package config for generated events
      packages:
        - path: github.com/foomo/sesamy-go/pkg/event
          types:
            - AddToCart
            - Purchase

# --- Conversion Linker settings
conversionLinker:
  # Enable provider
  enabled: true
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: ad_storage

# --- Umami settings
umami:
  # Enable provider
  enabled: true
  # Enter an optional fixed domain to override event data
  domain: your-domain.com
  # Paste ID for your website from the Umami settings
  websiteId: ''
  # Endpoint url of the umami api
  endpointUrl: https://umami.your-domain.com
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: analytics_storage
  # Google Tag Manager server container settings
  serverContainer:
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - PageView
          - SelectItem

# --- Facebook
# https://developers.facebook.com/docs/marketing-api/conversions-api/guides/gtm-server-side
facebook:
  # Enable provider
  enabled: true
  # Facebook pixel id
  pixelId: ''
  # To use the Conversions API, you need an access token.
  apiAccessToken: ''
  # Code used to verify that your server events are received correctly by Conversions API
  testEventToken: ''
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: ad_storage
  # Google Tag Manager server container settings
  serverContainer:
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - AddPaymentInfo
          - AddToCart
          - AddToWishlist
          - PageView
          - Purchase
          - Search
          - BeginCheckout
          - GenerateLead
          - ViewItem

# --- Emarsys
emarsys:
  # Enable provider
  enabled: true
  # Emarsys merchant id
  merchantId: ''
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: analytics_storage
  # Google Tag Manager server container settings
  serverContainer:
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - Purchase
          - ViewItem
          - ViewItemList

# --- Tracify
tracify:
  # Enable provider
  enabled: true
  # Tracify token
  token: ''
  # Tracify customer site id
  customerSiteId: ''
  # Google Consent settings
  googleConsent:
    # Enable consent mode
    enabled: true
    # Consent mode name
    mode: analytics_storage
  # Google Tag Manager server container settings
  serverContainer:
    # Contemplate package config for generated events
    packages:
      - path: github.com/foomo/sesamy-go/pkg/event
        types:
          - AddToCart
          - PageView
          - ViewItem
          - Purchase

# --- Cookiebot CMP
cookiebot:
  # Enable provider
  enabled: true
  # Name of the manually installed Cookiebot CMP tag template
  # "https://tagmanager.google.com/gallery/#/owners/cybotcorp/templates/gtm-templates-cookiebot-cmp
  templateName: Cookiebot CMP
  # Cookiebot id
  cookiebotId: ''
  # CDN Region (eu, com)
  cdnRegion: eu
  # Enable URL passthrough
  urlPassthrough: false
  # Enable advertiser consent mode
  advertiserConsentModeEnabled: false
```

## Caveats

You might need to increase your Google Tag Manager API quotas, since they are limited to 15 r/m by default.

## How to Contribute

Make a pull request...

## License

Distributed under MIT License, please see license file within the code for more details.
