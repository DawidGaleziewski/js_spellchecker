import React from 'react';
import {css} from '@emotion/react';
import styled from '@emotion/styled';

export enum ButtonVariants {
    PRIMARY = "primary",
    SECONDARY = "secondary",
    FORM_PRIMARY ="form-primary",
    FORM_SECONDARY = "form-secondary"
}

export interface IButtonProps {
    variant: ButtonVariants
}

export const Button = styled.button<IButtonProps>`
    cursor: pointer;
    padding: 1.6rem 3.7rem;
    border: none;
    font-size: ${props => props.theme.typography.buttonDefault.size};
    font-weight: ${props => props.theme.typography.buttonDefault.weight};
    font-family: ${props => props.theme.typography.buttonDefault.family};


    ${props => props.variant === ButtonVariants.PRIMARY && css`
        background-color: ${props.theme.color.primary};
        color: ${props.theme.color.textInverted};
        opacity: 1;
        transition: 0.3s linear opacity;
        outline: 0;
        border-radius: ${props.theme.radius[700]};

        &:focus-visible {
            outline: 1px solid ${props.theme.color.status.active};
            outline-offset: 2px;
        }

        &:hover, &:active {
            opacity: 0.8;
            outline: none;
        }

        &:disabled {
            background-color: ${props.theme.color.status.disabled};
            cursor: not-allowed;
            opacity: 1;
        }
    `}

    ${props => props.variant === ButtonVariants.SECONDARY && css`
        border: 1px solid ${props.theme.color.backgroundLight};
        background-color: ${props.theme.color.transparent};
        color: ${props.theme.color.backgroundLight};
        transition: background-color 0.3s linear;
        border-radius: ${props.theme.radius[700]};

        &:focus-visible {
            outline: 1px solid ${props.theme.color.status.active};
            outline-offset: 2px;
        }

        &:hover, &:active {
            background-color: ${props.theme.color.backgroundLight};
            border-color: ${props.theme.color.backgroundLight};
            color: ${props.theme.color.text};
        }

        &:disabled {
            background-color: ${props.theme.color.status.disabled};
            cursor: not-allowed;
            opacity: 1;
        }
    `}

    ${props => props.variant === ButtonVariants.FORM_PRIMARY && css`
        border-radius: ${props.theme.radius[100]};
        background-color: ${props.theme.color.primary};
        color: ${props.theme.color.textInverted};
        opacity: 1;
        transition: 0.3s linear opacity;
        outline: 0;

        &:focus-visible {
            outline: 1px solid ${props.theme.color.status.active};
            outline-offset: 2px;
        }

        &:hover, &:active {
            opacity: 0.8;
            outline: none;
        }

        &:disabled {
            background-color: ${props.theme.color.status.disabled};
            cursor: not-allowed;
            opacity: 1;
        }
    `}

    ${props => props.variant === ButtonVariants.FORM_SECONDARY && css`
        border-radius: ${props.theme.radius[100]};
        background-color: ${props.theme.color.secondary4};
        color: ${props.theme.color.textInverted};
        opacity: 1;
        transition: 0.3s linear background-color;
        outline: 0;

        &:focus-visible {
            outline: 1px solid ${props.theme.color.status.active};
            outline-offset: 2px;
        }

        &:hover, &:active {
            background-color: ${props.theme.color.status.active};
            outline: none;
        }

        &:disabled {
            background-color: ${props.theme.color.status.disabled};
            cursor: not-allowed;
            opacity: 1;
        }
    `}
`;


