## Release Notes Creation Guidance:
### Version number
- (must) consist of three numbers separated by dots
```
The leftmost number (1) is called the major version.
The middle number (2) is called the minor version.
The rightmost number (3) is called the revision but it may also be referred to as a "point release" or "subminor version".
```

### Release Notes file:
- use markdown format
- releases notes for appcenter app: `/release_notes/dev/`
- releases notes for production app: `/release_notes/production/`
- group files by minor version: e.g: 0.25.md, 0.26.md
- contains release notes of subminor versions:

#### File Content Format:
- new version with:
```
## VERSION: {version_number}
```

- if the release notes is use for both iOS & Android
```
## VERSION: {version_number}
{ReleaseNotes}
```

- if the release notes is different for iOS & Android
```
## VERSION: {version_number}
#### [iOS]
{iOS ReleaseNotes}

#### [Android]
{android ReleaseNotes}

```
- if the version is only available on iOS or android
```
## VERSION: {version_number}
#### [iOS]
{iOS ReleaseNotes}
```

<hr/>

### Sample:
## VERSION: 0.25.5
### New features:
- Fix the memory issue for large collection. Known issues: the artworks images take more time to load, it doesn't crop into the center when displaying in the home page.
- Enable Ledger for all built variants. Please try connecting Ledger from an Android phone.

### Updates:
- [UI] Fix arrow alignment and LINKED icon on Receive NFT UI

## VERSION: 0.25.6
#### [iOS]
- This is iOS
- Use Cloudflare to speed up the image loading. 

#### [Android]
- This is Android
- Use Cloudflare IPFS to see if it speeds up the artwork preview from IPFS files.

## VERSION: 0.25.7
#### [Android]
- This is changelogs for Android


## VERSION: 0.25.8
#### [iOS]
- This is changelogs for iOS

### Result:

![Simulator Screen Shot - iPhone 11 Pro - 2022-03-25 at 19 08 52](https://user-images.githubusercontent.com/50816149/160119712-145f775e-b2d1-460d-8b0d-2389534229e8.png)

