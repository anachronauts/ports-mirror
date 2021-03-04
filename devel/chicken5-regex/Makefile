# $FreeBSD$

PORTNAME=	regex
DISTVERSION=	2.0
PORTREVISION=	0
CATEGORIES=	devel
MASTER_SITES=	https://code.call-cc.org/egg-tarballs/5/${PORTNAME}/
PKGNAMEPREFIX=	chicken5-

MAINTAINER=	sage@anachronauts.club
COMMENT=	Old regex API for CHICKEN

RUN_DEPENDS=	chicken5:lang/chicken5
BUILD_DEPENDS=	chicken-install5:lang/chicken5

CSC_OPTIONS=	-v -strip

do-build:
	cd ${WRKSRC} && chicken-install5 \
		-v -no-install -no-install-dependencies

do-install:
	cd ${WRKSRC} && DESTDIR=${STAGEDIR} chicken-install5 \
		-v -no-install-dependencies

.export CSC_OPTIONS
.include <bsd.port.mk>
