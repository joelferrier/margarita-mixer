########
Commands
########

*************
General Usage
*************

.. code-block:: text

  margarita-mixer help
  margarita-mixer is LiME module build tool which
  uses docker to compile LiME kernel modules and creates
  a searchable repository of kernel modules.
  
  Usage:
    margarita-mixer [command]
  
  Available Commands:
    build       build LiME kernel modules
    help        Help about any command
    profile     perform operations on profiles
    version     Print the version number of margarita-mixer
  
  Flags:
    -h, --help   help for margarita-mixer
  
  Use "margarita-mixer [command] --help" for more information about a command.


*****
build
*****

.. code-block:: text

  margarita-mixer build
  build LiME kernel modules
  
  Usage:
    margarita-mixer build [command]
  
  Available Commands:
    all         build all configured profiles
  
  Flags:
    -h, --help   help for build
  
  Use "margarita-mixer build [command] --help" for more information about a command.

build all
=========

.. code-block:: text

  margarita-mixer build all
  project opened
  in newClient
  in configure:
  in setup
  in images
  sha256:b05c3d76c8b3ef3af8974edda3941a4a028d244681cc420e5304a3829519f1b6
  ...
  sha256:6133b2c7d7c2fb402770857b7ebaa1eff3613c21cc257aac42a8de49e3c7f74e
  in build
  in extract
  in cleanup

********
profiles
********

.. code-block:: text

  margarita-mixer profile help
  perform operations on profiles
  
  Usage:
    margarita-mixer profiles [command]
  
  Available Commands:
    list        list availible profiles
  
  Flags:
    -h, --help   help for profiles
  
  Use "margarita-mixer profile [command] --help" for more information about a command.

profiles list
=============

.. code-block:: text

  margarita-mixer profile list
  
  profiles:
  
  [amzn1]
    image: amazonlinux:1
    package manager: yum
    header packages: [kernel-devel.x86_64]
    dependencies: [git gcc make]
  
  [amzn2]
    image: amazonlinux:2
    package manager: yum
    header packages: [kernel-devel.x86_64]
    dependencies: [git gcc make]
  
  [centos6]
    image: centos:6
    package manager: yum
    header packages: [kernel-devel.x86_64]
    dependencies: [git gcc make]
  ...

*******
version
*******

.. code-block:: text

  margarita-mixer version
  margarita-mixer v0.1.0
