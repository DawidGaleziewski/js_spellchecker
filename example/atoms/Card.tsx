import React, { ReactElement } from "react";
import { Button, ButtonVariants, IButtonProps } from "./Button";
import { Typography, TypographyVariant, ITypography } from "./Typography";


enum CardVariants {
    LIGHT = "lingth"
}

interface ICardProps{
    variant: CardVariants;
    children: [React.ReactElement<ITypography>, React.ReactElement<IButtonProps>];
}

export const Card = (props: ICardProps) => <div>{props.children}</div>

const CardButton = (props:IButtonProps) => <Button variant={ButtonVariants.SECONDARY}>{props.children}</Button>;

Card.Button = CardButton;

const CardTitle:React.VFC<ITypography> = (props) => <Typography>{props.children}</Typography>;

Card.Title = CardTitle;

const CardTest = () => {
    return <Card variant={CardVariants.LIGHT}>
       <Card.Title variant={TypographyVariant.LINK}>Test</Card.Title>
       <Card.Button variant={ButtonVariants.PRIMARY}>Test</Card.Button>
    </Card>
}