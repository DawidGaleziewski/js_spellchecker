import React from 'react';

import styled from '@emotion/styled';
import { css } from "@emotion/react";



export const Table = (props) => <div>{props.children}</div>;

const TableRowBase = styled.div`
   display: flex;
`;
const TableRow = (props) => <TableRowBase>{props.children}</TableRowBase>;
Table.Row = TableRow;


const TableCellBase = styled.div`
    padding: 1.6rem;
`;
const TableCell = (props) => <TableCellBase>{props.children}</TableCellBase>;
Table.Cell = TableCell;

