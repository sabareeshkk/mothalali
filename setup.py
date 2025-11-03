#!/usr/bin/env python3

from setuptools import setup, find_packages

setup(
    name='mothalali',
    version='1.0.0',
    packages=['mothalali'],
    entry_points={
        'console_scripts': [
            'mothalali=mothalali.cli:main',
        ],
    },
    author='sabareesh k',
    author_email='sabareeshk1991@gmail.com',
    description='Mothalali Git is a toy git project onspired from ugit',
    url='https://github.com/mothalali/mothalali-git',
    classifiers=[
        'Programming Language :: Python :: 3',
    ],
    python_requires='>=3.6',
)