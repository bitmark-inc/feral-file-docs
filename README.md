# Feral File Documents Repository

## Overview
This repository serves as a **maintenance document** for managing and structuring various documentation files used across Feral File applications, including the mobile app, website, and TV display app. The files in this repository provide essential agreements, configurations, and content necessary for different features of the platform.

## Impact on Production
- Any changes made to existing files that are merged into the `main` branch will **directly affect the app and web in the PRD environment**.
- **Developers must be informed when adding new files** so they can include them in the data loading process for the app and website.

## Repository Structure

### **1. agreements/**
Contains legal and policy-related documents that are loaded as static pages on the app and website. These files include:
- app data usage
- app eula
- app privacy
- art transfer agreement
- artist collector rights
- artist collector term of service
- web privacy
- ... (additional legal agreements)

### **2. app/**
Holds various data files used by the **Feral File app**, including:
- `release-note.md` - Release notes for the app.
- `security.md` - Security-related documentation.
- **Source exhibition data** - Contains structured data for exhibitions displayed in the app.
- `version.json` - Tracks the versioning details of the app.

### **3. configs/**
This folder contains JSON configuration files that are used by:
- **Mobile App**
- **TV Display App**

These files define app settings and behaviors for different environments.

### **4. docs/**
This is a legacy folder created earlier in the repository’s lifecycle, and it **cannot be renamed** due to system dependencies. It primarily stores:
- **Templates for artist and curator contracts**

The templates are structured within `docs/curator-artist-agreement/`, which generate contracts for each exhibition.

### **5. learning/**
Content for **white papers** and educational materials. Each Markdown file in this directory corresponds to an article in the **Learn section** of the Feral File website.
Developers must be informed when adding new files so the article can be load with SEO metadata in the website.

## Maintenance Notes
- **agreements/**: Updates to policies or agreements should be reviewed before deployment.
- **app/**: Only the Feral File app dev team is working on this.
- **configs/**: Changes here affect mobile and TV applications, requiring careful testing.
- **docs/**: Since this directory is tied to contract generation, modifications should be handled with caution.
- **learning/**: New white papers should follow the established Markdown formatting before being published.

For any structural changes, ensure that dependencies across apps and websites are considered before modifying this repository.
