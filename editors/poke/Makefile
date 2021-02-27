# $FreeBSD$

PORTNAME=	poke
DISTVERSION=	1.0
CATEGORIES=	editors
MASTER_SITES=	GNU

MAINTAINER=	sage@anachronauts.club
COMMENT=	Interactive, extensible editor for binary data

USES=		gmake libtool pkgconfig:build

LIB_DEPENDS=	libgc.so:devel/boehm-gc \
		libreadline.so:devel/readline \
		libjson-c.so:devel/json-c \
		libtextstyle.so:devel/libtextstyle

GNU_CONFIGURE=	yes
CONFIGURE_ARGS=	--disable-gui --disable-libnbd

USE_LDCONFIG=	yes

.include <bsd.port.mk>
