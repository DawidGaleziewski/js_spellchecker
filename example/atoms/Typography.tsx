import React from 'react';
import { css } from '@emotion/react';
import styled from '@emotion/styled';

export const TypographyVariant = {
    MAIN_HEADER: "main-header",
    SECTION_HEADER: "section-header",
    SUB_HEADER: "sub-header",
    CARD_HEADER:"card-header",
    BUTTON_TEXT: 'button-text',
    MOBILE_MENU: "mobile-menu",
    PARAGRAPH: "paragraph",
    LIST: "list",
    LINK: "link"
} as const;

type TypographyVariantUnionType = typeof TypographyVariant[keyof typeof TypographyVariant]

const VariantDefaultHTMLTag = {
    [TypographyVariant.MAIN_HEADER]: 'h1',
    [TypographyVariant.SECTION_HEADER]: 'h2',
    [TypographyVariant.SUB_HEADER]: 'h3',
    [TypographyVariant.CARD_HEADER]: 'h5',
    [TypographyVariant.BUTTON_TEXT]: 'span',
    [TypographyVariant.MOBILE_MENU]: 'span',
    [TypographyVariant.PARAGRAPH]: 'p',
    [TypographyVariant.LIST]: 'li',
    [TypographyVariant.LINK]: 'span'
}

export interface ITypography {
    variant?: TypographyVariantUnionType;
    as?: React.ElementType<any>;
    children?: React.ReactElement
}

export const Typography: React.VFC<ITypography> = (props = { variant: TypographyVariant.PARAGRAPH}) => <TypographyBase variant={props.variant} as={props.as ?? VariantDefaultHTMLTag[props.variant ?? TypographyVariant.PARAGRAPH] }>
    {props.children}
</TypographyBase>;


const TypographyBase = styled.p<ITypography>`
    margin: 0;
    ${props => props.variant === TypographyVariant.MAIN_HEADER && css`
        font-size: ${props.theme.typography.mainHeader.size};
        font-weight: ${props.theme.typography.mainHeader.weight};
        font-family: ${props.theme.typography.mainHeader.family};
        line-height: ${props.theme.typography.mainHeader.height};
    `};

    ${props => props.variant === TypographyVariant.SECTION_HEADER && css`
        font-size: ${props.theme.typography.sectionHeader.size};
        font-weight: ${props.theme.typography.sectionHeader.weight};
        font-family: ${props.theme.typography.sectionHeader.family};
        line-height: ${props.theme.typography.sectionHeader.height};
    `};

    ${props => props.variant === TypographyVariant.SUB_HEADER && css`
        font-size: ${props.theme.typography.subHeader1.size};
        font-weight: ${props.theme.typography.subHeader1.weight};
        font-family: ${props.theme.typography.subHeader1.family};
        line-height: ${props.theme.typography.subHeader1.height};
    `};

    ${props => props.variant === TypographyVariant.CARD_HEADER && css`
        font-size: ${props.theme.typography.h5.size};
        font-weight: ${props.theme.typography.h5.weight};
        font-family: ${props.theme.typography.h5.family};
        line-height: ${props.theme.typography.h5.height};
    `};

    ${props => props.variant === TypographyVariant.BUTTON_TEXT && css`
        font-size: ${props.theme.typography.buttonDefault.size};
        font-weight: ${props.theme.typography.buttonDefault.weight};
        font-family: ${props.theme.typography.buttonDefault.family};
        line-height: ${props.theme.typography.buttonDefault.height};
    `};

    ${props => props.variant === TypographyVariant.PARAGRAPH && css`
        font-size: ${props.theme.typography.paragraph.size};
        font-weight: ${props.theme.typography.paragraph.weight};
        font-family: ${props.theme.typography.paragraph.family};
        line-height: ${props.theme.typography.paragraph.height};
    `}
`;