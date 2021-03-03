# $FreeBSD$

PORTNAME=	srfi-14
DISTVERSION=	1.4
PORTREVISION=	0
CATEGORIES=	devel
MASTER_SITES=	https://code.call-cc.org/egg-tarballs/5/${PORTNAME}/
PKGNAMEPREFIX=	chicken5-

MAINTAINER=	sage@anachronauts.club
COMMENT=	Character set library

RUN_DEPENDS=	chicken5:lang/chicken5
BUILD_DEPENDS=	chicken-install5:lang/chicken5

CSC_OPTIONS=	-v -strip
C5_REPO=	${:!chicken-install5 -repository!}

do-build:
	cd ${WRKSRC} && chicken-install5 \
		-v -no-install -no-install-dependencies

do-install:
	cd ${WRKSRC} && DESTDIR=${STAGEDIR} chicken-install5 \
		-v -no-install-dependencies

.export CSC_OPTIONS
.include <bsd.port.mk>
