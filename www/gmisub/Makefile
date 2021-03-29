# $FreeBSD$

PORTNAME=	gmisub
DISTVERSION=	g20210220
PORTREVISION=	0
CATEGORIES=	www

MAINTAINER=	sage@anachronauts.club
COMMENT=	Subscribe to Gemini pages

MASTER_SITES=	https://git.sr.ht/~callum/gmisub/archive/
DISTFILES=	b84a2381.tar.gz
WRKSRC=		${WRKDIR}/${PORTNAME}-b84a2381

USES=		gmake
BUILD_DEPENDS=	scdoc:textproc/scdoc
LIB_DEPENDS=	libgmni.so:www/gmni

.include <bsd.port.mk>
