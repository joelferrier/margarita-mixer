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
    help        Help about any command
    profile     perform operations on profiles
    version     Print the version number of margarita-mixer
  
  Flags:
    -h, --help   help for margarita-mixer
  
  Use "margarita-mixer [command] --help" for more information about a command.

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
