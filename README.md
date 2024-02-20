# Overview
LDS is a CLI tool for generating a Software Bill of Materials (SBOM) from container images and filesystems. The generated SBOM file is well comply with the SPDX v2.3 specification.
The SBOM can be used for vulnerability detection through scanner like [Wind River Scanning Tool](https://studio.windriver.com/scan).

---------------------------------------------------------------------------------------
## Supporting systems
- Ubuntu
- Debian
- Red Hat / CentOs

## Supporting scan targets
- Docker image
- Filesystem

## LDS Output
LDS generate SBOM for the target system. The SBOM collects the system packages, the dependency and License info.

----------------------------------------------------------------------------------------
# Quick Start
## Get LDS

### Download released LDS binary
Download and extract LDS of your architecture from https://github.com/Wind-River/wr-linux-distro-sbom/releases

### Build LDS from source code
Require golang version >= 1.21

```bash
git clone https://github.com/Wind-River/wr-linux-distro-sbom
cd wr-linux-distro-sbom/cmd/lds
go build
```


## Generate SBOM
LDS generate SBOM in SPDX v2.3 format.

To generate SBOM for a container image:
```bash
./lds packages <image:tag> --file <output-filename.spdx.json>
```

To generate SBOM for a container image archive file:
```bash
./lds packages <image-archive-name.tar> --file <output-filename.spdx.json>
```

To generate SBOM for a directory:
```bash
./lds packages <filesystem-directory> --file <output-filename.spdx.json>
```

---------------------------------------------------------------------------------------
# Legal Notices

All product names, logos, and brands are property of their respective owners.

All company, product and service names used in this software are for identification purposes only.

Wind River is a trademark of Wind River Systems, Inc.

Disclaimer of Warranty / No Support: Wind River does not provide support and maintenance services for this software,

under Wind River’s standard Software Support and Maintenance Agreement or otherwise.

Unless required by applicable law, Wind River provides the software (and each contributor

provides its contribution) on an “AS IS” BASIS, WITHOUT WARRANTIES OF ANY

KIND, either express or implied, including, without limitation, any warranties

of TITLE, NONINFRINGEMENT, MERCHANTABILITY, or FITNESS FOR A PARTICULAR

PURPOSE. You are solely responsible for determining the appropriateness of

using or redistributing the software and assume any risks associated with

your exercisa of permissions under the license.
