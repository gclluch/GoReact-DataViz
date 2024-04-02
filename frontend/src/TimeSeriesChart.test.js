import React from 'react';
import axios from 'axios';
import { render, screen, waitFor } from '@testing-library/react';
import TimeSeriesChart from './TimeSeriesChart';

jest.mock('axios');

// Ensures the TimeSeriesChart component fetches data using axios, processes the data
// and attempts to render a chart by checking for the presence of a canvas element
test('TimeSeriesChart fetches and displays data correctly', async () => {
    axios.get.mockResolvedValue({ data: [{ x: 1, y: 1 }, { x: 2, y: 2 }, { x: 3, y: 3 }] });

    render(<TimeSeriesChart url="http://localhost:8080/api/series1" title="Test Chart" />);

    await waitFor(() => {
        const canvasElement = screen.getByRole('img');
        expect(canvasElement).toBeInTheDocument();
    });
});
