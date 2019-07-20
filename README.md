# on-server

Web file server for download and upload files. Userful for quick file sharing.

### Installation

Locate to on-server folder and run 
```bash
go build -o ./bin/on-server
```
command to build module.

### Usage

Execute ```./bin/on-server``` file, open given address with web browser. :on:

Command-line flags list.

* ``-ip``    Server ip address, if not provided listens to all interfaces
*  ``-port`` Port number (default "2100")

* ``-path``  Server files root (default current path)
* ``-message-path`` Text message files location (default current path)
* ``-upload-path`` Uploaded files location (default current path)

* ``-no-files`` Disable files listing
* ``-no-message`` Disable text submit
* ``-no-upload`` Disable file upload

* ``-send-limit`` Maximum allowed data to send in MegaBytes (default 64)
* ``-show-ip6`` Show IP6 addresses in list if 'ip' not provided
        
