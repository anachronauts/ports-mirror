# $FreeBSD$

PORTNAME=	srfi-18
DISTVERSION=	1.4
PORTREVISION=	0
CATEGORIES=	devel
MASTER_SITES=	https://code.call-cc.org/egg-tarballs/5/${PORTNAME}/
PKGNAMEPREFIX=	chicken5-

MAINTAINER=	sage@anachronauts.club
COMMENT=	A multithreading package, largely following the specification of SRFI-18

RUN_DEPENDS=	chicken5:lang/chicken5
BUILD_DEPENDS=	chicken-install5:lang/chicken5

CSC_OPTIONS=	-v -strip
C5_REPO:=	${:!chicken-install5 -repository!}
CHICKEN_INSTALL_REPOSITORY=	${STAGEDIR}${C5_REPO}

do-build:
	cd ${WRKSRC} && chicken-install5 \
		-v -no-install -no-install-dependencies

do-install:
	cd ${WRKSRC} && chicken-install5 -v -no-install-dependencies


.export CSC_OPTIONS CHICKEN_INSTALL_REPOSITORY
.include <bsd.port.mk>
