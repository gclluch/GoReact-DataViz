import React from 'react';
import { Line } from 'react-chartjs-2';
import {
    Chart as ChartJS,
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend,
} from 'chart.js';
import useFetchData from './ChartDataFetcher';

ChartJS.register(
    CategoryScale,
    LinearScale,
    PointElement,
    LineElement,
    Title,
    Tooltip,
    Legend
);

const TimeSeriesChart = ({ url, title }) => {
    const dataPoints = useFetchData(url);

    if (dataPoints.length === 0) {
        return <div>Loading...</div>;
    }

    const data = {
        labels: dataPoints.map(point => point.x),
        datasets: [
        {
            label: title,
            data: dataPoints.map(point => point.y),
            borderColor: 'rgb(75, 192, 192)',
            backgroundColor: 'rgba(75, 192, 192, 0.5)',
        },
        ],
    };

    const options = {
        scales: {
        y: {
            beginAtZero: true
        },
        x: {
            type: 'linear',
            position: 'bottom'
        }
        },
        responsive: true,
        plugins: {
        legend: {
            position: 'top',
        },
        title: {
            display: true,
            text: title,
        },
        },
    };

    return <Line options={options} data={data} />;
};

export default TimeSeriesChart;
